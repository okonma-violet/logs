package logger

import (
	"time"
)

// const ()

// const TagsMaxLength = 65535

// func (lf *logframe) encodeLog() []byte {
// 	log := encode(tags, logstr, name)
// 	log[0] = logtype.Byte()
// 	if !logtime.IsZero() {
// 		byteOrder.PutUint64(log[1:], uint64(logtime.UnixMicro()))
// 	}
// 	return log
// }

// func xAppendTags(tags []byte, newtags ...string) []byte {
// 	if len(tags) == 0 {
// 		return encode(make([]byte, 11), "", newtags...)
// 	}
// 	return encode(tags, "", newtags...)
// }

// func encode(tags []byte, logstr string, newtags ...string) []byte {
// 	var tgs []byte
// 	if len(tags) == 0 {
// 		tgs = make([]byte, 11)
// 	} else {
// 		tgs = tags
// 	}
// 	tagslen := len(tgs)
// 	tslist := make([][]byte, 0, len(newtags))
// 	for _, tg := range newtags {
// 		if len(tg) > 0 {
// 			tb := make([]byte, len(tg)+3)

// 			n := copy(tb[1:], tg)
// 			tb[0] = TagStartSep
// 			tb[n+1] = TagEndSep
// 			tb[n+2] = TagDelim

// 			tagslen += len(tb)
// 			tslist = append(tslist, tb)
// 		}
// 	}
// 	if tagslen > TagsMaxLength {
// 		panic("tags length is out of range (lib logs/encode), last tag:" + newtags[len(newtags)-1])
// 	}
// 	log := make([]byte, len(tgs), tagslen+len(logstr))
// 	copy(log, tgs)
// 	for _, t := range tslist {
// 		log = append(log, t...)
// 	}
// 	log = append(log, logstr...)
// 	byteOrder.PutUint16(log[9:], uint16(tagslen))
// 	return log
// }

// func DecodeToString(log []byte) string {
// 	if len(log) < 11 {
// 		panic("logs/encode/DecodeToString() recieved log with len less than 11") // TODO:?
// 	}
// 	return suckutils.Concat(string(TagStartSep), LogType(log[0]).String(), string(TagEndSep), time.UnixMicro(int64(byteOrder.Uint64(log[1:9]))).Format(time_layout), string(log[11:]))
// }
// func DecodeToStringColorized(log []byte) string {
// 	if len(log) < 11 {
// 		panic("logs/encode/DecodeToString() recieved log with len less than 11") // TODO:?
// 	}
// 	return suckutils.Concat(LogType(log[0]).Colorize(), string(TagStartSep), LogType(log[0]).String(), string(TagEndSep), ColorWhite, time.UnixMicro(int64(byteOrder.Uint64(log[1:9]))).Format(time_layout), string(log[11:]))
// }

func PrintLn(logtype LogType, tag string, text string) {
	println(concatLocalLogFrameBody(logtype, time.Now(), nil, tag, text))
	//println(concatLocalLogFrameBody(logtype, time.Now(), appendLocalTags(nil, tags), "", text))
}
