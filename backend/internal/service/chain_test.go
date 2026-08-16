package service
import (
 "errors"
 "testing"
 "gbevent/internal/repository"
)
func TestErrorChainSentinels(t *testing.T){
 db:=newEvDB(t)
 rr:=repository.NewRegistrationRepository(db); ar:=repository.NewActivityRepository(db); nr:=repository.NewNotificationRepository(db)
 if _,err:=rr.FindByID(999);!errors.Is(err,repository.ErrNotFound){t.Fatalf("reg FindByID = %v, want ErrNotFound",err)}
 if _,err:=ar.FindByID(999);!errors.Is(err,repository.ErrNotFound){t.Fatalf("act FindByID = %v, want ErrNotFound",err)}
 if _,err:=nr.FindByID(999);!errors.Is(err,repository.ErrNotFound){t.Fatalf("notif FindByID = %v, want ErrNotFound",err)}
}
