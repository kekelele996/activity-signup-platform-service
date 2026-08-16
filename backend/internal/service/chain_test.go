package service
import (
 "testing"
 "time"
 "gbevent/internal/constants"
 "gbevent/internal/repository"
)
func TestSignupCapacityAndDedup(t *testing.T){
 db:=newEvDB(t)
 ar:=repository.NewActivityRepository(db); rr:=repository.NewRegistrationRepository(db)
 cr:=repository.NewCheckInRecordRepository(db); nr:=repository.NewNotificationRepository(db)
 as:=NewActivityService(ar,rr,nr,cr,evLogger()); rs:=NewRegistrationService(db,rr,as,nr,evLogger())
 now:=time.Now()
 a,err:=as.Create(1,"讲座","","","lecture","A",now.Add(time.Hour),now.Add(2*time.Hour),now.Add(30*time.Minute),2,constants.ActivityStatusPublished)
 if err!=nil{t.Fatalf("Create activity: %v",err)}
 for i,u:=range []uint64{11,12}{ if _,err:=rs.Create(a.ID,u,"n","p","");err!=nil{t.Fatalf("signup %d: %v",i,err)} }
 if _,err:=rs.Create(a.ID,13,"n","p","");err==nil{t.Fatal("third signup should be full")}
 if _,err:=rs.Create(a.ID,11,"n","p","");err==nil{t.Fatal("duplicate signup should fail")}
}
