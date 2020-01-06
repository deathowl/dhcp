// +build freebsd openbsd netbsd darwin

package client4
import 	(
	"syscall"

	"github.com/insomniacslk/dhcp/dhcpv4"
	"golang.org/x/sys/unix"

)

func makeListeningSocketWithCustomPort(ifname string, port int) (int, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	if err != nil {
		return fd, err
	}
	err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
	if err != nil {
		return fd, err
	}

    llAddr := unix.SockaddrInet4{Addr: [4]byte{127, 0, 0, 1}}
	err = dhcpv4.BindToInterface(fd, ifname)
	if err != nil {
		return fd, err
	}

	err = unix.Bind(fd, &llAddr)
	return fd, err

}


// makeRawSocket creates a socket that can be passed to unix.Sendto.
func makeRawSocket(ifname string) (int, error) {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_RAW, unix.IPPROTO_RAW)
	if err != nil {
		return fd, err
	}
	err = unix.SetsockoptInt(fd, unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
	if err != nil {
		return fd, err
	}
	err = unix.SetsockoptInt(fd, unix.IPPROTO_IP, unix.IP_HDRINCL, 1)
	if err != nil {
		return fd, err
	}
	err = dhcpv4.BindToInterface(fd, ifname)
	if err != nil {
		return fd, err
	}
	return fd, nil
}
