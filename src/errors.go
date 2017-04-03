package main

type CouldNotOpenFileError struct {
	message string
}

func NewCouldNotOpenFileError(msg string) *CouldNotOpenFileError {
	return &CouldNotOpenFileError{
		message: msg}
}

func (CouldNotOpenFileError *CouldNotOpenFileError) Error() string {
	return CouldNotOpenFileError.message
}

///////
type ScanInputFileError struct {
	message string
}

func NewScanInputFileError(msg string) *ScanInputFileError {
	return &ScanInputFileError{
		message: msg}
}

func (scanInputFileError *ScanInputFileError) Error() string {
	return scanInputFileError.message
}

//////
type ScanValueError struct {
	message string
}

func NewScanValueError(msg string) *ScanValueError {
	return &ScanValueError{
		message: msg}
}

func (scanValueError *ScanValueError) Error() string {
	return scanValueError.message
}
