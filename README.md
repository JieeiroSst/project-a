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