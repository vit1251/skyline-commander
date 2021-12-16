package tty

//#include <sys/select.h>
//#define STDIN 0
//int wait_input() {
//  fd_set fds;
//  int maxfd;
//  maxfd = STDIN;
//  FD_ZERO(&fds);
//  FD_SET(STDIN, &fds);
//  select(maxfd+1, &fds, NULL, NULL, NULL);
//  if (FD_ISSET(STDIN, &fds)){
//      return 1;
//  }
//  return 0;
//}
import "C"
import "fmt"

func waitInput() error {
	err1 := C.wait_input()
	if err1 != 0 {
		return fmt.Errorf("wrong await result on select")
	}
	return nil
}
