package wxxx

import "fmt"

type Logger interface {
	PrintLn()
	Printf()
}

type logger struct {
}

const logTag = "[WXXX]"

//noinspection ALL
func (*logger) PrintLn(a ...interface{}) {
	vWithTimeAndTag := make([]interface{}, 0, len(a)+2)
	vWithTimeAndTag = append(vWithTimeAndTag, TimeNow().Format(longTimeWithMsFormat), logTag)
	vWithTimeAndTag = append(vWithTimeAndTag, a...)
	fmt.Println(vWithTimeAndTag...)
}

//noinspection ALL
func (*logger) Printf(format string, a ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s %s %s\n", TimeNow().Format(longTimeWithMsFormat), logTag, format), a...)
}
