package service
import (
 "testing"
 "gbevent/internal/model"
 "gbevent/internal/repository"
)
func TestNotificationChain(t *testing.T){
 db:=newEvDB(t)
 nr:=repository.NewNotificationRepository(db)
 n1:=&model.Notification{UserID:11,NotificationType:"signup_success",Title:"t",Content:"c"}; db.Create(n1)
 n2:=&model.Notification{UserID:12,NotificationType:"signup_success",Title:"t2",Content:"c2"}; db.Create(n2)
 if err:=nr.MarkRead(n1.ID,12);err!=nil{t.Fatalf("MarkRead: %v",err)}
 got,_,_:=nr.ListByUser(11,1,10)
 if got[0].IsRead{t.Fatal("user11 notification should not be read by user12")}
 if got[0].ID!=n1.ID{t.Fatal("list order wrong")}
}
