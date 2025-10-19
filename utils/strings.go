package utils

import (
	"bytes"
	cryptoRand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unsafe"

	"net/mail"

	"github.com/dongri/phonenumber"
	"github.com/google/uuid"
	"github.com/yunarsuanto/base-go/constants"

	log "github.com/sirupsen/logrus"

	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	idCode = "+62"
	male   = "M"
	female = "F"
)

var (
	phoneNumberRegex         = regexp.MustCompile("^[+]?[0-9]{8,}$")
	emailRegex               = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	civilPlateRegex          = regexp.MustCompile(`^([A-Z]{1,3})(\s|-)*([1-9][0-9]{0,4})(\s|-)*([A-Z]{0,3}|[1-9][0-9]{1,2})$`)
	militaryPolicePlateRegex = regexp.MustCompile(`^([0-9]{1,5})(\s|-)*([0-9]{2}|[IVX]{1,5})$`)
	letterBytes              = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	numberBytes              = "0123456789"
)

func GenerateOtp() string {
	var letterIdxBits int64 = 6                    // 6 bits to represent a letter index
	var letterIdxMask int64 = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	var letterIdxMax int64 = 10 / letterIdxBits    // # of letter indices fitting in 63 bits
	var src = rand.NewSource(time.Now().UnixNano())

	b := make([]byte, constants.DefaultChangePasswordOtpLength)
	for i, cache, remain := constants.DefaultChangePasswordOtpLength-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(numberBytes) {
			b[i] = numberBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// ValidateEmail return true if email address has a valid format
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// HashPassword function to hash password
// Params:
// password: password to hash
// Returns hashed string
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash function to check if hashed string and password is compatible
// Params:
// password: password
// hash: hashed string
// Returns bool
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// NormalizePhoneNumber standarized mobile phone number format
func NormalizePhoneNumber(number string) string {
	number = strings.Replace(strings.TrimSpace(number), " ", "", -1)

	if len(number) < 8 {
		return number
	}

	if strings.HasPrefix(number, "0") {
		return fmt.Sprintf("%s%s", idCode, number[1:])
	}

	if strings.HasPrefix(number, "+62") {
		return fmt.Sprintf("%s%s", idCode, number[3:])
	}

	origNumber := number
	for {
		if len(number) <= 2 || !strings.HasPrefix(number, "62") {
			break
		}

		number = number[2:]
	}

	if len(number) < 8 {
		number = origNumber
	}

	return fmt.Sprintf("%s%s", idCode, number)
}

// ValidatePhoneNumber return true if number has a valid phone number format
func ValidatePhoneNumber(number string) bool {
	number = NormalizePhoneNumber(number)
	return phoneNumberRegex.MatchString(number)
}

func trimQuote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

// DateString return string representation of a date pointer
func DateString(dateTime *time.Time) string {
	if dateTime == nil {
		return ""
	}

	return dateTime.Format(time.DateOnly)
}

// NullBoolScanToString function to convert bool pointer to string
// Params:
// x: bool pointer
// Returns string
// if x is null, returns "-"
func NullBoolScanToString(x *bool) string {
	if x != nil {
		return strconv.FormatBool(*x)
	}
	return "-"
}

// ValidateBloodType function to validate blood type input
// Params:
// bloodType: blood type to validate
// Returns *constants.ErrorResponse
func ValidateBloodType(bloodType string) *constants.ErrorResponse {
	validBloodType := []string{
		constants.APlusBloodType,
		constants.BPlusBloodType,
		constants.ABPlusBloodType,
		constants.OPlusBloodType,
		constants.AMinusBloodType,
		constants.BMinusBloodType,
		constants.ABMinusBloodType,
		constants.OMinusBloodType,
	}

	for _, v := range validBloodType {
		if v == bloodType {
			return nil
		}
	}

	return constants.ErrInvalidBloodType
}

// ConvertBytesToString function to convert byte array to string
// Params:
// data: byte array
// Returns string
func ConvertBytesToString(data []byte) string {
	return string(data[:])
}

// StructToByte function to convert struct to array byte
// Params:
// data: struct to convert
// Returns array byte and *constants.ErrorResponse
func StructToByte(data any) ([]byte, *constants.ErrorResponse) {
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(data)
	if err != nil {
		return nil, ErrorInternalServer(err.Error())
	}

	return reqBodyBytes.Bytes(), nil
}

// HourMinuteToTimeFormat function to convert hour and minute to time format
// Params:
// hour: hour
// minute: minute
// Returns time-formatted string
func HourMinuteToTimeFormat(hour, minute int) string {
	hourString := fmt.Sprintf("00%d", hour)
	minuteString := fmt.Sprintf("00%d", minute)

	hourStringFormatted := hourString[len(hourString)-2:]
	minuteStringFormatted := minuteString[len(minuteString)-2:]

	return fmt.Sprintf("%s:%s", hourStringFormatted, minuteStringFormatted)
}

// TimeFormatToHourMinute function to convert time-formatted string to hour and minute
// Params:
// s: time formatted string
// Returns hour, minute, and *constants.ErrorResponse
func TimeFormatToHourMinute(s string) (int, int, *constants.ErrorResponse) {
	timeArray := strings.Split(s, ":")
	hour, err := strconv.Atoi(timeArray[0])
	if err != nil {
		return 0, 0, ErrorInternalServer(err.Error())
	}
	minute, err := strconv.Atoi(timeArray[1])
	if err != nil {
		return 0, 0, ErrorInternalServer(err.Error())
	}

	return hour, minute, nil
}

// GetUserTitle function to get user title according to gender and marital status
// Params:
// gender: gender, valid value: M, F
// isMarried: bool, if user is married
// Returns title string and *constants.ErrorResponse
func GetUserTitle(gender string, isMarried bool) (string, *constants.ErrorResponse) {
	if gender == male {
		return "Mr", nil
	}
	if gender == female {
		if isMarried {
			return "Mrs", nil
		}
		return "Ms", nil
	}

	return "", constants.ErrInvalidGender
}

// SplitUserName function to
// Params:
// Returns
func SplitUserName(name string) (string, string, *constants.ErrorResponse) {
	name = strings.TrimSpace(name)

	nameArray := strings.Split(name, " ")
	var firstName string
	var lastName string

	if len(nameArray) == 0 {
		return firstName, lastName, constants.ErrUserNameAbsence
	}
	if len(nameArray) == 1 {
		firstName = nameArray[0]
		lastName = nameArray[0]
	} else {
		firstName = strings.Join(nameArray[:len(nameArray)-1], " ")
		lastName = nameArray[len(nameArray)-1:][0]
	}

	return firstName, lastName, nil
}

// PrintStruct function to print struct in json sytled format
// Params:
// data: struct to be printed in log
func PrintStruct(data any) {
	dataString := StructToJson(data)
	log.Println(dataString)
}

// StructToJson function to convert struct to json
// Params:
// data: struct to be converted
func StructToJson(data any) string {
	dataByte, errs := StructToByte(data)
	if errs != nil {
		log.Errorln(errs.Err)
	}

	return string(dataByte)
}

// PrintError function to print error in json sytled format
// Params:
// err: error to be printed in log
func PrintError(err constants.ErrorResponse) {
	results := struct {
		HttpCode int
		Message  string
	}{
		HttpCode: err.HttpCode,
		Message:  err.Err.Error(),
	}

	dataByte, errs := StructToByte(results)
	if errs != nil {
		log.Errorln(errs.Err)
	}
	log.Errorln(string(dataByte))
}

// Uid function to generate uid
// Params:
// length: string length
// Returns uid
func Uid(length int) string {
	buf := []string{}
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-="
	charRune := []rune(chars)
	charlen := len(chars)

	for i := 0; i < length; i++ {
		index := rand.Intn(charlen)
		buf = append(buf, string(charRune[index]))
	}

	return strings.Join(buf, "")
}

// ConvertMapToString function to convert map[string]string to url params styled string
// Params:
// m: value to be converted
// Returns url params styled string
func ConvertMapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\" ", key, value)
	}
	return b.String()
}

// IsLetter function to check whether string only contains alphabetic character
// Params:
// s: string to be checked
// Returns boolean
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// NewNullString function to handle nullable value in SQL
// Params:
// s: string to be checked
// Returns sql.NullString
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// ValidateEmailOrPhoneNumberInput function to handle email or phone number in single string
// Params:
// input: string to be validated
// Returns isEmail bool, isPhoneNumber bool
func ValidateEmailOrPhoneNumberInput(input *string) (bool, bool) {
	var isEmail bool
	var isPhoneNumber bool

	inputAsPhoneNumber := phonenumber.Parse(*input, constants.IndonesiaAlphaCode)
	if inputAsPhoneNumber != "" {
		*input = inputAsPhoneNumber
		isPhoneNumber = true
	} else {
		_, err := mail.ParseAddress(*input)
		if err != nil {
			return isEmail, isPhoneNumber
		}
		isEmail = true
	}

	return isEmail, isPhoneNumber
}

// QueryLog function for log query on repository
func QueryLog(query string, args ...any) {
	mappingVar := []map[string]string{}
	for i, v := range args {
		mappingVar = append(mappingVar, map[string]string{
			"key":   fmt.Sprintf("$%d", (i + 1)),
			"value": fmt.Sprintf("'%v'", v),
		})
	}

	for _, mapping := range reverseArray(mappingVar) {
		query = strings.ReplaceAll(query, mapping["key"], mapping["value"])
	}
	log.Println(NormalizeString(query))
}

func NormalizeString(s string) string {
	result := strings.ReplaceAll(s, "\n", " ")
	result = strings.ReplaceAll(result, "\t", " ")
	result = strings.Join(strings.Fields(result), " ")

	return result
}

// QueryLogNamed function for log query on repository
func QueryLogNamed(query string, input any) {
	stringTag := strings.Split(query, ":")
	stringTags := []string{}
	mappingVar := []map[string]string{}
	if len(stringTag) > 0 {
		for i := 0; i < len(stringTag); i++ {
			regexCompile := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
			tagString := regexCompile.ReplaceAllString(stringTag[i], "")
			tagString = regexCompile.ReplaceAllString(tagString, " ")
			tagString = strings.ReplaceAll(tagString, ",", "")
			tagString = strings.ReplaceAll(tagString, "\n", "")
			tagString = strings.ReplaceAll(tagString, ")", "")
			if i != 0 {
				stringTags = append(stringTags, tagString)
			}
		}
	}
	rt := reflect.TypeOf(input)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		resTag, ok := f.Tag.Lookup("db")
		if !ok {
			return
		}
		if InArrayExist(resTag, stringTags) {
			val := reflect.ValueOf(input).FieldByName(f.Name)
			mappingVar = append(mappingVar, map[string]string{
				"key":   fmt.Sprintf(":%s", resTag),
				"value": fmt.Sprintf("'%v'", val),
			})
		}
	}
	for _, mapping := range reverseArray(mappingVar) {
		query = strings.ReplaceAll(query, mapping["key"], mapping["value"])
	}
	log.Println(NormalizeString(query))
}

// QueryLog QueryLogMysql
func QueryLogMysql(query string, args ...any) {
	query = strings.ReplaceAll(query, "?", "'%v'")
	log.Println(NormalizeString(fmt.Sprintf(query, args...)))
}

func reverseArray(input []map[string]string) []map[string]string {
	if len(input) == 0 {
		return input
	}
	return append(reverseArray(input[1:]), input[0])
}

// GenerateRedisKey function to normalize redis key
func GenerateRedisKey(appName, id, purpose, platform, uniqueKey string, allowMultipleLogin bool) string {
	var result string
	if allowMultipleLogin {
		result = fmt.Sprintf("%s-%s-%s:%s-%s", appName, purpose, platform, id, uniqueKey)
	} else {
		result = fmt.Sprintf("%s-%s-%s:%s", appName, purpose, platform, id)
	}
	return result
}

// RandomString Generate Random String
func RandomString(n int) string {
	var letterIdxBits int64 = 6                    // 6 bits to represent a letter index
	var letterIdxMask int64 = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	var letterIdxMax int64 = 63 / letterIdxBits    // # of letter indices fitting in 63 bits
	var src = rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// SafetyString function to handle pointer string nulled
// Params:
// s: string to be checked
// Returns string
func SafetyString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// SafetyFloat function to handle pointer string nulled
// Params:
// s: float64 to be checked
// Returns float64
func SafetyFloat(s *float64) float64 {
	if s == nil {
		return 0
	}
	return *s
}

// SafetyNullSqlString function to handle pointer string nulled
// Params:
// s: sql.NullString to be checked
// Returns string
// Returns string
func SafetyNullSqlString(s sql.NullString) string {
	if !s.Valid {
		return ""
	}
	return s.String
}

// StringToDateTime function to handle pointer string nulled
// Params:
// s: string to be checked
// Returns time.Time
// Returns January 1, year 1, 00:00:00 UTC. if empty string
func StringToDateTime(s string) (time.Time, *constants.ErrorResponse) {
	var t time.Time
	if s == "" {
		return t, nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return t, ErrorInternalServer(err.Error())
	}
	return t, nil
}

// StringToDate function to handle pointer string nulled
// Params:
// s: string to be checked
// Returns time.Time
// Returns January 1, year 1, 00:00:00 UTC. if empty string
func StringToDate(s string) (time.Time, *constants.ErrorResponse) {
	var t time.Time
	if s == "" {
		return t, nil
	}
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return t, ErrorInternalServer(err.Error())
	}
	return t, nil
}

// StringToTime function to handle pointer string nulled
// Params:
// s: string to be checked
// Returns time.Time
// Returns January 1, year 1, 00:00:00 UTC. if empty string
func StringToTime(s string) (time.Time, *constants.ErrorResponse) {
	var t time.Time
	if s == "" {
		return t, nil
	}
	t, err := time.Parse(constants.HourMinuteOnly, s)
	if err != nil {
		return t, ErrorInternalServer(err.Error())
	}
	t = t.AddDate(1, 0, 0)
	return t, nil
}

// StringToFloat64 function to handle pointer string nulled
// Params:
// s: string to be checked
// Returns float64
// Returns 0 if empty string
func StringToFloat64(s string) float64 {
	if s == "" {
		return 0
	}
	f, e := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if e != nil {
		return 0
	}
	return f
}

var months = [...]string{
	"Januari",
	"Februari",
	"Maret",
	"April",
	"Mei",
	"Juni",
	"Juli",
	"Agustus",
	"September",
	"Oktober",
	"November",
	"Desember",
}
var days = [...]string{
	"Minggu",
	"Senin",
	"Selasa",
	"Rabu",
	"Kamis",
	"Jumat",
	"Sabtu",
}

// MonthIndonesia return string representation of a date pointer
func MonthIndonesia(dateTime time.Time, short bool) string {
	m, _ := strconv.Atoi(dateTime.Format("1"))
	if short {
		return months[m-1][:3]
	}
	return months[m-1]
}

// DayIndonesia return string representation of a date pointer
func DayIndonesia(dateTime time.Time) string {
	d := int(dateTime.Weekday())
	return days[d]
}

func IsValidUUID(u string) *constants.ErrorResponse {
	_, err := uuid.Parse(u)
	if err != nil {
		return ErrorInternalServer(err.Error())
	}
	return nil
}

func CensorName(s string) string {
	strArr := strings.Split(s, " ")

	res := []string{}
	for _, v := range strArr {
		stringLength := len(v)
		halfLength := stringLength / 2
		restLength := stringLength - halfLength

		newString := v[:halfLength]
		for i := 0; i < restLength; i++ {
			newString = fmt.Sprintf("%s%s", newString, "*")
		}
		res = append(res, newString)
	}

	return strings.Join(res, " ")
}

func GenerateRsaPair(length int) (privateKey, publicKey string, err error) {
	privateKeyBase, err := rsa.GenerateKey(cryptoRand.Reader, length)
	if err != nil {
		return privateKey, publicKey, err
	}
	privatePem := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKeyBase)}
	privatePemBytes := pem.EncodeToMemory(privatePem)
	privateKey = string(privatePemBytes)

	publicKeyBase := &privateKeyBase.PublicKey
	publicPem := &pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(publicKeyBase)}
	publicPemBytes := pem.EncodeToMemory(publicPem)
	publicKey = string(publicPemBytes)

	return privateKey, publicKey, err
}

func CompareRsaPair(privateKey, publicKey string) (matched bool, err error) {
	block, _ := pem.Decode([]byte(privateKey))
	loadedPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return matched, err
	}

	block, _ = pem.Decode([]byte(publicKey))
	loadedPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return matched, err
	}

	matched = loadedPrivateKey.PublicKey.Equal(loadedPublicKey)

	return matched, err
}

func StandarizeSex(s string) string {
	if s == constants.IndonesianMale {
		return constants.UserSexMale
	}
	if s == constants.IndonesianFemale {
		return constants.UserSexFemale
	}

	return s
}

func StringToBoolPointer(s string) (*bool, *constants.ErrorResponse) {
	if s == "" {
		return nil, nil
	}

	res, err := strconv.ParseBool(s)
	if err != nil {
		return nil, ErrorInternalServer(err.Error())
	}

	return &res, nil
}

func FormatVehicleNumber(input string) (string, *constants.ErrorResponse) {
	s := strings.ToUpper(strings.TrimSpace(input))

	if civilPlateRegex.MatchString(s) {
		out := civilPlateRegex.ReplaceAllString(s, `$1 $3 $5`)
		return strings.TrimSpace(out), nil
	}

	if militaryPolicePlateRegex.MatchString(s) {
		out := militaryPolicePlateRegex.ReplaceAllString(s, `$1-$3`)
		return strings.TrimSpace(out), nil
	}

	return "", constants.ErrInvalidVehicleNumber
}
