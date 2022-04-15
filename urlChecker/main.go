package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

//for channel
type resultRequest struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {

	// var을 사용하여 초기화 방법
	var results = make(map[string]string)
	results2 := make(map[string]string)
	ch := make(chan resultRequest)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	//results["gello"] = "Hello"

	//before checking urls concurrently through goroutines
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "Failed"
		}
		results[url] = result
	}
	//fmt.Println(results)

	for url, result := range results {
		fmt.Println(url, result)
	}

	//after checking urls concurrently through goroutines
	for _, url := range urls {
		go hitURLfast(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		result := <-ch
		results2[result.url] = result.status
	}

	for url, result := range results {
		fmt.Println(url, result)
	}

	// concurrently running
	go sexyCount("nico")
	go sexyCount("flynn")
	//maintain function for seconds
	time.Sleep(time.Second * 8)

	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}

	// isSexy 메소드로 실행될때 채널을 통해해메세지 동시에옴
	// waiting for message is blocking the operation
	// goroutine이 없는데 메세지를 기다리면 실행을 block하기 떄문에 프로그램이 hang걸려 멈춘다
	//fmt.Println("Received: ", <-c)
	//fmt.Println("Received: ", <-c)
	// 받을 메세지가 없는데 기다리면 deadlock
	//fmt.Println(<-c)

	for i := 0; i < len(people); i++ {
		fmt.Println("Waiting for: ", i)
		fmt.Println("Received: ", <-c)
	}

	// ** rules of channel **
	// 먼저해야하는걸 해치운다
	// 메인 함수가 끝나면 고루틴은 죽는다 무조건
	// 어떤타입의 데이터를 보내고 받을지 채널을 통해서 지정해준다
	// 메세지가 받을수 없는 곳도 보낼수 있다. 받는곳은 따로지정하기 때문에
	// 메시지를 받는곳을 지정하면 실행을 블락한다
}

func hitURL(url string) error {
	fmt.Println("checking: " + url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}

//send-only channel
func hitURLfast(url string, ch chan<- resultRequest) {
	fmt.Println("checking: " + url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	ch <- resultRequest{url: url, status: status}
}

// goroutine은 function이 실행되는동안만 유지된다.
// ex) 함수가 부르는 모든 메소드에 go를 붙이면 그냥 종료됨
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

// 채널 생성, 메세지 리턴
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	//fmt.Println(person)
	c <- person + " is sexy"
}
