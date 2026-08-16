package service
import (
 "testing"
 "time"
 "gbevent/internal/constants"
 "gbevent/internal/repository"
)
func TestCheckinAndStats(t *testing.T){
 db:=newEvDB(t)
 ar:=repository.NewActivityRepository(db); rr:=repository.NewRegistrationRepository(db)
 cr:=repository.NewCheckInRecordRepository(db); nr:=repository.NewNotificationRepository(db)
 as:=NewActivityService(ar,rr,nr,cr,evLogger()); rs:=NewRegistrationService(db,rr,as,nr,evLogger())
 cs:=NewCheckInRecordService(db,cr,rr,as,nr,evLogger())
 now:=time.Now()
 a,_:=as.Create(1,"讲座","","","lecture","A",now.Add(time.Hour),now.Add(2*time.Hour),now.Add(30*time.Minute),10,constants.ActivityStatusPublished)
 reg,err:=rs.Create(a.ID,11,"n","p","")
 if err!=nil{t.Fatalf("signup: %v",err)}
 if _,err:=cs.CheckInByVoucher(a.ID,1,reg.VoucherNo);err!=nil{t.Fatalf("checkin: %v",err)}
 st,err:=as.Stats(a.ID,1,constants.RoleAdmin)
 if err!=nil{t.Fatalf("stats: %v",err)}
 if st["checkin_rate"]!=100.0{t.Fatalf("checkin_rate = %v, want 100",st["checkin_rate"])}
}
