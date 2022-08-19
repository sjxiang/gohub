
// 实现 "gorm.io/gorm/logger" Interface，代替它，塞点私货进去

package logger


// GormLogger 操作对象
type GormLogger struct {

}


/*

2022/08/20 01:36:38 /home/xsj/go/src/github.com/sjxiang/gohub/app/data/user/user_util.go:18
[0.336ms] [rows:1] SELECT count(*) FROM `users` WHERE phone = '18018001800'

*/
