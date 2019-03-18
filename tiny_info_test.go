package Figo

import (
	"github.com/quexer/red"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestRedTinyInfo_Put(t *testing.T) {
	pool := red.CreatePool(10, "127.0.0.1:6379", "")
	redDo := red.BuildDoFunc(pool)

	tf := NewTinyInfo(redDo, "test")
	v, err := tf.Put("hello")
	logrus.WithField("seq", v).WithField("err", err).Println()
	v, err = tf.Put("world")
	logrus.WithField("seq", v).WithField("err", err).Println()
	v, err = tf.Put("how")
	logrus.WithField("seq", v).WithField("err", err).Println()
	v, err = tf.Put("r")
	logrus.WithField("seq", v).WithField("err", err).Println()
	v, err = tf.Put("u")
	logrus.WithField("seq", v).WithField("err", err).Println()
}

func TestRedTinyInfo_Get(t *testing.T) {
	pool := red.CreatePool(10, "127.0.0.1:6379", "")
	redDo := red.BuildDoFunc(pool)

	tf := NewTinyInfo(redDo, "test")
	for i := 1; i < 5; i++ {
		content, _ := tf.Get(i)
		logrus.WithField("val", content).Println()
	}
}
