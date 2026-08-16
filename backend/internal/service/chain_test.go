package service
import (
 "errors"
 "testing"
 "gbevent/internal/constants"
 "gbevent/internal/repository"
)
func expectNoPanicEv(t *testing.T, name string, fn func() error) error { t.Helper(); var err error; func(){ defer func(){ if r:=recover(); r!=nil { t.Fatalf("%s panicked: %v",name,r) } }(); err=fn() }(); return err }
func TestMissingRegistrationAndActivity(t *testing.T){
 db:=newEvDB(t)
 rr:=repository.NewRegistrationRepository(db); ar:=repository.NewActivityRepository(db)
 cr:=repository.NewCheckInRecordRepository(db); nr:=repository.NewNotificationRepository(db)
 as:=NewActivityService(ar,rr,nr,cr,evLogger()); rs:=NewRegistrationService(db,rr,as,nr,evLogger())
 err:=expectNoPanicEv(t,"Cancel",func() error { _,err:=rs.Cancel(999,1,constants.RoleAdmin); return err })
 if !errors.Is(err,repository.ErrNotFound){t.Fatalf("Cancel = %v, want ErrNotFound",err)}
 err=expectNoPanicEv(t,"Stats",func() error { _,err:=as.Stats(999,1,constants.RoleAdmin); return err })
 if !errors.Is(err,repository.ErrNotFound){t.Fatalf("Stats = %v, want ErrNotFound",err)}
}
