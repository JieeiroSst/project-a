mockgen -destination=mock_doer.go -package=repository github.com/JieeiroSst/itjob/users/internal/repository UserRepository
 
mockgen -destination=mocks/mock_doer.go -package=mocks github.com/JieeiroSst/itjob/users/internal/usecase UserUsecase
 
mockgen -destination=mocks/mock_doer.go -package=mocks github.com/JieeiroSst/itjob/users/internal/http UserHttp
 
mockgen -destination=mocks/mock_doer.go -package=mocks  github.com/JieeiroSst/itjob/users/internal/delivery/http DeliveryHttp
 
mockgen -destination=mocks/mock_doer.go -package=mocks  github.com/JieeiroSst/itjob/users/internal/db UserDB
 
github.com/JieeiroSst/itjob/users/internal
 
 
 
.Preload("Orders", func(db *gorm.DB) *gorm.DB {
   return db.Unscoped() 
}

func Monit(db *gorm.DB) {
	log.Println("[INFO][System]\tStarted monitoring of files and db entries")
	tc := time.NewTicker(1 * time.Minute)
	for {
		res := []models.ResourceEntry{}
		db.Find(&res, "created_at < ?", time.Now().Add(-timeLimit))
		db.Unscoped().Where("created_at < ?", time.Now().Add(-timeLimit)).Delete(&models.ResourceEntry{})
		if len(res) > 0 {
			log.Printf("[INFO][System]\tFlushing %d DB entries and files.\n", len(res))
		}
		for _, re := range res {
			err := os.Remove(path.Join(conf.C.UploadDir, re.Key))
			if err != nil {
				log.Printf("[ERROR][System]\tWhile deleting : %v", err)
			}
		}
		<-tc.C
	}
}

func (p *Point) Delete(db gorm.DB) {
	db.Unscoped().Delete(p)
}

time_start _+ interval '10 second'


func unique(sample []Data) []Data {
	var unique []Data
sampleLoop:
	for _, v := range sample {
		for i, u := range unique {
			if v.factoryId == u.factoryId && v.plantId == u.plantId && v.lineId == u.lineId {
				unique[i] = v
				continue sampleLoop
			}
		}
		unique = append(unique, v)
	}
	return unique
}


open premission send email . https://myaccount.google.com/u/0/lesssecureapps


. call rabbitmq
```
func (e *EmailServer) Run() {
	emailsPublisher, err := rabbitmq.NewEmailsPublisher(e.cfg)
	if err != nil {
		log.Println("Emails Publisher can't initialized")
	}
	defer emailsPublisher.CloseChan()

	err = emailsPublisher.SetupExchangeAndQueue(
		e.cfg.RabbitMQ.Exchange,
		e.cfg.RabbitMQ.Queue,
		e.cfg.RabbitMQ.RoutingKey,
		e.cfg.RabbitMQ.ConsumerTag, )
	if err!=nil{
		log.Println(err)
	}

	repo:=repository.NewEmailRepository(e.db,*e.snowflake)
	usecase:=usecase.NewEmailUseCase(repo,emailsPublisher,*e.utils,*e.emailPkg)
	http.NewHttp(e.server,usecase)
	emailsAmqpConsumer := rabbitmq.NewImagesConsumer(e.amqpConn, usecase,e.cfg)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := emailsAmqpConsumer.StartConsumer(
			e.cfg.RabbitMQ.WorkerPoolSize,
			e.cfg.RabbitMQ.Exchange,
			e.cfg.RabbitMQ.Queue,
			e.cfg.RabbitMQ.RoutingKey,
			e.cfg.RabbitMQ.ConsumerTag,
		)
		if err != nil {
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		log.Printf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Printf("ctx.Done: %v", done)
	}
}
```


```
app.Get("/:id", func(c *fiber.Ctx) error {
    id := c.Params("id")
    res, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + id)
    if err != nil {
        return err
    }

    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return err
    }

    todo := Todo{}
    parseErr := json.Unmarshal(body, &todo)
    if parseErr != nil {
        return parseErr
    }

    return c.JSON(fiber.Map{"Data": todo})
})
```

```
package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/ReneKroon/ttlcache/v2"
    "github.com/gofiber/fiber/v2"
)

type Todo struct {
    Userid    int    `json:"userId"`
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

var cache ttlcache.SimpleCache = ttlcache.NewCache()

func verifyCache(c *fiber.Ctx) error {
    id := c.Params("id")
    val, err := cache.Get(id)
    if err != ttlcache.ErrNotFound {
        return c.JSON(fiber.Map{"Cached": val})
    }
    return c.Next()
}

func main() {
    app := fiber.New()

    cache.SetTTL(time.Duration(10 * time.Second))

    app.Get("/:id", verifyCache, func(c *fiber.Ctx) error {
        id := c.Params("id")
        res, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + id)
        if err != nil {
            return err
        }

        defer res.Body.Close()
        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
            return err
        }

        todo := Todo{}
        parseErr := json.Unmarshal(body, &todo)
        if parseErr != nil {
            return parseErr
        }

        cache.Set(id, todo)
        return c.JSON(fiber.Map{"Data": todo})
    })

    app.Listen(":3000")
}

```