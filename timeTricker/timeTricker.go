package main

import (
    "fmt"
//    "sync"
//    "time"
//	"gitlab.xunlei.cn/xlppc/redpacket/model/mysql_db"
	"gopkg.in/robfig/cron.v2"
//	"sync"
//	"time"
)

// func main() {
// 	hour := time.Now()
// 	fmt.Println(hour)

// 	fmt.Println(hour.Hour())
// }

// func task() {
// 	var wg sync.WaitGroup
// 	//mysql_db.GetVisitCount(0, 200)
// 	pipe := make(chan mysql_db.PacketsPool, 200)

// 	wg.Add(1)
// 	go func(group *sync.WaitGroup) {
// 		FromLastDay(pipe)
// 		fmt.Println("FromLastDay exec suc")
// 		group.Done()
// 	}(&wg)

// 	wg.Add(1)
// 	go func(group *sync.WaitGroup) {
// 		ToTomorrow(pipe)
// 		fmt.Println("ToTomorrow exec suc")
// 		group.Done()
// 	}(&wg)
// 	wg.Wait()
// }

// func FromLastDay(pipe chan mysql_db.PacketsPool) {
// 	timeNow := time.Now()
// 	dataTime := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())
// 	retS, err := mysql_db.GetLastDayPacket(dataTime, 100)
// 	if err != nil {
// 		fmt.Println("GetLastDayPacket err: ", err)
// 		goto EXIT
// 	}
// 	for _, ret := range retS {
// 		pipe <- ret
// 	}
// EXIT:
// 	// 关闭 chan 后 for range 读取完数据后会退出循环, 否则一直等待
// 	close(pipe)
// }

// func ToTomorrow(ch chan mysql_db.PacketsPool) {
// 	dataTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
// 	for v := range ch {
// 		v.Id = 0
// 		v.DateId = v.DateId.AddDate(0, 0, 1)
// 		if v.StartTime.Before(dataTime) {
// 			continue
// 		}
// 		v.StartTime = v.StartTime.AddDate(0, 0, 1)
// 		ret, _ := mysql_db.CreatePacketToPool(v)
// 		fmt.Println("affect rows = ", ret)
// 	}
// }

func start() {
	c := cron.New()
	a := "1"
	t := fmt.Sprintf(" %s * * * * *", a)
	fmt.Println(t)
	c.AddFunc(t, func() {
		fmt.Println("Every hour on the half hour")
		task()
		fmt.Printf("task finish\n")
	})
    c.Start()
    //select {}
}


func main() {
    fmt.Printf("start")
    // c := cron.New()
	// c.AddFunc("1 * * * * *", func() {
	// 	fmt.Println("Every hour on the half hour")
	// 	task()
	// 	fmt.Printf("task finish\n")
	// })
    // c.Start()
    // select {}
    //go start()

    // todayTime := time.Now().UnixNano() // 10e6
    // fmt.Println("time:", todayTime)

    i := 0
    for ; i < 10; i++ {
        fmt.Println(i)
    }
    select {}
    fmt.Printf("end\n")
}


func task() {
    fmt.Println("hello")
}



// func main() {
//     w := sync.WaitGroup{}
//     i := 0
//     for ; i < 10 ; i++ {
//         w.Add(1)
//         go func(b int) {
//             fmt.Println(b)
//             w.Done()
//         }(i)
//     }
//     w.Wait()
// }





// package main

// import (
//     "time"
//     "fmt"
// )

// //定时结算Boottime表数据
// func BoottimeTimingSettlement() {
//     for {
//         now := time.Now()
//         // 计算下一个零点
//         next := now.Add(time.Hour * 3)
//         next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
//         t := time.NewTimer(next.Sub(now))
//         <-t.C
//         // fmt.Println("定时结算Boottime表数据，结算完成: %v\n",time.Now())
//         //以下为定时执行的操作
//         BoottimeSettlement()
//     }
// }

// func BoottimeSettlement() {
// 	fmt.Println("hello")
// 	fmt.Println("结算完成:",time.Now().Format("2019-02-26 00：00:00"))
// }

// func main() {
// 	fmt.Println("satrt")
// 	BoottimeTimingSettlement()
// 	fmt.Println("end")
// }
