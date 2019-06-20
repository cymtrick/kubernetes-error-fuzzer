// +build !selinux !linux

package label

// InitLabels returns the process label and file labels to be used within
// the container.  A list of options can be passed into this function to alter
// the labels.
func InitLabels(options []string) (string, string, error) {
	return "", "", nil
}

func ROMountLabel() string {
	return ""
}

func GenLabels(options string) (string, string, error) {
	return "", "", nil
}

func FormatMountLabel(src string, mountLabel string) string {
	return src
}

func SetProcessLabel(processLabel string) error {
	return nil
}

func ProcessLabel() (string, error) {
	return "", nil
}

func SetSocketLabel(processLabel string) error {
	return nil
}

func SocketLabel() (string, error) {
	return "", nil
}

func SetKeyLabel(processLabel string) error {
	return nil
}

func KeyLabel() (string, error) {
	return "", nil
}

func FileLabel(path string) (string, error) {
	return "", nil
}

func SetFileLabel(path string, fileLabel string) error {
	return nil
}

func SetFileCreateLabel(fileLabel string) error {
	return nil
}

func Relabel(path string, fileLabel string, shared bool) error {
	return nil
}

func PidLabel(pid int) (string, error) {
	return "", nil
}

func Init() {
}

// ClearLabels clears all reserved labels
func ClearLabels() {
	return
}

func ReserveLabel(label string) error {
	return nil
}

func ReleaseLabel(label string) error {
	return nil
}

// DupSecOpt takes a process label and returns security options that
// can be used to set duplicate labels on future container processes
func DupSecOpt(src string) ([]string, error) {
	return nil, nil
}

// DisableSecOpt returns a security opt that can disable labeling
// support for future container processes
func DisableSecOpt() []string {
	return nil
}

// Validate checks that the label does not include unexpected options
func Validate(label string) error {
	return nil
}

// RelabelNeeded checks whether the user requested a relabel
func RelabelNeeded(label string) bool {
	return false
}

// IsShared checks that the label includes a "shared" mark
func IsShared(label string) bool {
	return false
}
