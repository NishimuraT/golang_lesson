package main
/*
1.エラーの伝播
2.
 */
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

}

// 1. エラーの伝播
func throwError(url string) (string, error) {
	_, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return "", nil
}

// 2. エラーの伝播（＋リトライ、＋情報付加）
// NASAの自己調査：ジェネシス：クラッシュ：パラシュート開かず：Gスイッチ失敗：誤ったリレー方向（grepで操作できる様に１行）
func fetch(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0;time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

// 3. 処理を進める事が不可能であれば、呼び出しもとがエラーを表示してプログラムを上手く停止させる
// 一般的に、その様な処理はmainパッケージに限定されるべき
// os.Exit(1)

// 4. エラーを記録して、おそらく制限された機能で処理を続ける
//func ping() {
//	if err := Ping();err != nil {
//		log.Printf("ping failed: %v; networking disabled", err)
//		fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled", err)
//	}
//}