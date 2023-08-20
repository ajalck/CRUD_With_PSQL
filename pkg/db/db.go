package db

import (
	"log"
	"sync"

	"github.com/ajalck/CRUD_With_PSQL/pkg/config"
	"github.com/ajalck/CRUD_With_PSQL/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func ConnectDB(ch, ch1, ch2 chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	c := <-ch
	configuration := c.(*config.Config)
	db, err := gorm.Open(postgres.Open(configuration.DB_Source), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	creatingEnumType := `	DO $$
						BEGIN
							IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname='status')
							THEN
							CREATE TYPE status AS ENUM('active','terminated','placed','overdue','dropped');
							END IF;
						END $$;`
	db.Exec(creatingEnumType)
	if err := db.AutoMigrate(&models.StudentPortal{}); err != nil {
		log.Fatalf(err.Error())
	}

	ch1 <- db
	ch2 <- db
}

func ConfigDB(ch1 chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	DB := <-ch1
	db := DB.(*gorm.DB)
	tx := db.Begin()
	checkNameLength := `DO $$
						BEGIN
							IF NOT EXISTS(SELECT 1 FROM information_schema.table_constraints WHERE constraint_name='check_name')
							THEN 
							ALTER TABLE student_portal ADD CONSTRAINT check_name CHECK(LENGTH(name)>=3);
							END IF;
						END $$;`

	result := tx.Exec(checkNameLength)
	if result.Error != nil {
		result.Rollback()
	}

	checkEmail := `DO $$
					BEGIN
						IF NOT EXISTS(SELECT 1 FROM information_schema.table_constraints WHERE constraint_name='check_email')
						THEN
						ALTER TABLE student_portal ADD CONSTRAINT check_email CHECK(email~'^[A-Za-z0-9_.-%]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$');
						END IF;
					END $$;`

	result = tx.Exec(checkEmail)
	if result.Error != nil {
		result.Rollback()
	}

	checkBatch := `DO $$
					BEGIN
						IF NOT EXISTS(SELECT 1 FROM information_schema.table_constraints WHERE constraint_name='check_batch')
						THEN
						ALTER TABLE student_portal ADD CONSTRAINT check_batch CHECK(batch~'^BCE[0-9]');
						END IF;
					END $$;`

	result = tx.Exec(checkBatch)
	if result.Error != nil {
		result.Rollback()
	}

	checkDomain := `DO $$
					BEGIN
						IF NOT EXISTS(SELECT 1 FROM information_schema.table_constraints WHERE constraint_name='check_domain')
						THEN
						ALTER TABLE student_portal ADD CONSTRAINT check_domain CHECK(LENGTH(domain)>=3);
						END IF;
					END $$;`

	if result = tx.Exec(checkDomain); result.Error != nil {
		result.Rollback()
	}

	checkWeek := `DO $$
					BEGIN
						IF NOT EXISTS(SELECT 1 FROM information_schema.table_constraints WHERE constraint_name='check_week')
						THEN
						ALTER TABLE student_portal ADD CONSTRAINT check_week CHECK(week<29);
						END IF;
					END $$;`

	if result = tx.Exec(checkWeek); result.Error != nil {
		result.Rollback()
	}

	checEndDate := `DO $$
					BEGIN 
						IF NOT EXISTS(SELECT 1 FROM information_schema.table_constraints WHERE constraint_name='check_end_date')
						THEN
						ALTER TABLE student_portal ADD CONSTRAINT check_end_date CHECK(end_date IS NOT NULL AND status <> 'active');
						END IF;
					END $$;`

	if result = tx.Exec(checEndDate); result.Error != nil {
		result.Rollback()
	}
	tx.Commit()

}
