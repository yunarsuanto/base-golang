package constants

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type ErrorResponse struct {
	HttpCode int
	Err      error
}

const (
	DbDuplicateConstraint  = "Duplicate entry"
	DbForeignKeyConstraint = "foreign key constraint fails"
	DbTableDoesNotExist    = "pq: relation \"%s\" does not exist"
	DbDeleteRestriction    = "ON DELETE RESTRICT"
	JWTInvalidType         = "key is of invalid type"
	JWTExpired             = "token is expired by"
	RedisNilValue          = "redis: nil"
	PGNoRows               = "sql: no rows in result set"
)

const (
	ErrDuplicate                      = "terdapat data duplikat pada sistem: %s"
	ErrForeignKey                     = "tidak terdapat data referensi: %s"
	ErrHttpClient                     = "error when trying to access %s"
	ErrInvalidDateFormat              = "invalid Date Format: %s"
	ErrDataNotFound                   = "data %s tidak ditemukan"
	ErrInvalidOrderKey                = "invalid Order Key: %s"
	ErrInvalidOrder                   = "invalid Order: %s"
	ErrRedis                          = "error redis: %s"
	ErrTooFrequentMail                = "mohon menunggu selama %s sebelum mencoba kembali"
	ErrInvalidConditionTriggerSection = "invalid condition trigger section: %s"
	ErrInvalidConditionGroupOperator  = "invalid condition group operator: %s"
	ErrInvalidConditionOperator       = "invalid condition operator: %s"
	ErrInvalidFileType                = "invalid file type: %s"
	ErrFilledDefaultInputType         = "input type %s cannot have default value"
	ErrInvalidInputType               = "invalid value %s for input type %s"
	ErrApprovedEntry                  = "verifikasi \"%s\" untuk formulir ini sudah dilakukan"
)

func Error(httpCode int, message string, args ...string) *ErrorResponse {
	if len(args) != 0 {
		message = fmt.Sprintf(message, strings.Join(args, ", "))
	}

	errs := ErrorResponse{
		HttpCode: httpCode,
		Err:      errors.New(message),
	}

	return &errs
}

var CustomDatabaseUniqueKeyErrorResponse map[string]*ErrorResponse = map[string]*ErrorResponse{}

var (
	ErrTokenIsRequired                      = Error(http.StatusUnauthorized, "token is required")
	ErrKeyIsNotInvalidType                  = Error(http.StatusInternalServerError, JWTInvalidType)
	ErrIneligibleAccess                     = Error(http.StatusForbidden, "anda tidak memiliki hak akses untuk fitur ini")
	ErrInsufficientAccessHospital           = Error(http.StatusForbidden, "anda tidak memiliki akses ke rumah sakit")
	ErrDeleteRestriction                    = Error(http.StatusBadRequest, "tidak dapat menghapus data ini karena sudah ada data lain yang menggunakannya")
	ErrTokenInvalid                         = Error(http.StatusUnauthorized, "token is invalid")
	ErrSsoUserInvalid                       = Error(http.StatusUnauthorized, "user SSO tidak valid")
	ErrTokenReplaced                        = Error(http.StatusUnauthorized, "silakan login kembali untuk melanjutkan proses")
	ErrEmailAndPasswordNotMatch             = Error(http.StatusUnauthorized, "email atau password tidak sesuai")
	ErrEmailIsExist                         = Error(http.StatusBadRequest, "email dan/atau username sudah digunakan")
	ErrEmployeeNumberIsExist                = Error(http.StatusBadRequest, "NIP sudah digunakan")
	ErrIdNumberIsExist                      = Error(http.StatusBadRequest, "nomor identitas sudah digunakan")
	ErrPhoneNumberIsExist                   = Error(http.StatusBadRequest, "nomor telepon sudah digunakan")
	ErrPasswordNotMatch                     = Error(http.StatusBadRequest, "password tidak sesuai")
	ErrInvalidBloodType                     = Error(http.StatusBadRequest, "invalid Blood Type")
	ErrInvalidGender                        = Error(http.StatusBadRequest, "user has invalid gender value")
	ErrUserNameAbsence                      = Error(http.StatusBadRequest, "user has no name")
	ErrGenerateToken                        = Error(http.StatusBadRequest, "err generate from jwt")
	ErrSetTokenToRedis                      = Error(http.StatusBadRequest, "err set token to redis")
	ErrInvalidClaims                        = Error(http.StatusInternalServerError, "invalid claims")
	ErrInvalidMimeType                      = Error(http.StatusBadRequest, "anda tidak dapat mengunggah file ini")
	ErrInvalidUserUpdate                    = Error(http.StatusForbidden, "anda tidak dapat mengubah data ini")
	ErrInactiveUser                         = Error(http.StatusForbidden, "data pengguna anda telah tidak aktif. Silakan hubungi admin untuk informasi lebih lanjut")
	ErrUnavailableOTP                       = Error(http.StatusForbidden, "fitur OTP tidak tersedia untuk pengguna ini")
	ErrInvalidOtp                           = Error(http.StatusBadRequest, "OTP tidak sesuai. Silakan meminta OTP baru")
	ErrInvalidCronExpression                = Error(http.StatusBadRequest, "format cron tidak sesuai")
	ErrEmptyConditionGroup                  = Error(http.StatusBadRequest, "group condition should always have at least 1 condition")
	ErrFilledParentSection                  = Error(http.StatusBadRequest, "section or matrix column with childs should only be 'none' type")
	ErrNonFileInputWithType                 = Error(http.StatusBadRequest, "cannot add file type for section or matrix column with input type other than 'file'")
	ErrStaticValueRequired                  = Error(http.StatusBadRequest, "teks statis diperlukan untuk tipe elemen 'text editor'")
	ErrEmptyEnum                            = Error(http.StatusBadRequest, "data input 'radio', 'select', 'checkbox', atau 'autocomplete' tidak boleh kosong")
	ErrEmptyFormVersionHospital             = Error(http.StatusBadRequest, "data rumah sakit tidak boleh kosong")
	ErrEmptyFormVersionRole                 = Error(http.StatusBadRequest, "data peran pengguna tidak boleh kosong")
	ErrEmptyFormVersionGroup                = Error(http.StatusBadRequest, "formulir tidak boleh kosong")
	ErrEmptyFormVersionApproval             = Error(http.StatusBadRequest, "persetujuan formulir tidak boleh kosong")
	ErrEmptyFormVersionApprovalRole         = Error(http.StatusBadRequest, "pejabat persetujuan formulir tidak boleh kosong")
	ErrEmptyFormVersionGroupSection         = Error(http.StatusBadRequest, "tidak dapat mengajukan formulir dengan section kosong")
	ErrEmptyMatrixSection                   = Error(http.StatusBadRequest, "section dengan tipe element matrix harus memiliki kolom dan baris")
	ErrInvalidFormVersion                   = Error(http.StatusBadRequest, "versi formulir tidak valid")
	ErrInvalidFormVersionStatusUpdate       = Error(http.StatusBadRequest, "versi formulir ini sudah tidak dapat diubah karena statusnya sudah bukan draft lagi")
	ErrInvalidFormVersionStatusApprove      = Error(http.StatusBadRequest, "formulir ini sudah direspon / belum diajukan")
	ErrInvalidFormVersionStatusEntry        = Error(http.StatusBadRequest, "versi formulir ini tidak dapat diisi karena statusnya belum disetujui")
	ErrInvalidFormVersionStatusReference    = Error(http.StatusBadRequest, "formulir referensi belum disetujui")
	ErrInvalidFormEntryStatusUpdate         = Error(http.StatusBadRequest, "isian formulir ini sudah tidak dapat diubah karena statusnya sudah bukan draft lagi")
	ErrInvalidFormEntryStatusApprove        = Error(http.StatusBadRequest, "formulir ini sudah direspon / belum diajukan")
	ErrMandatoryFormEntryReference          = Error(http.StatusBadRequest, "dibutuhkan formulir referensi untuk mengisi formulir ini")
	ErrInvalidFormReference                 = Error(http.StatusBadRequest, "formulir referensi tidak valid")
	ErrInvalidHospitalId                    = Error(http.StatusForbidden, "anda tidak memiliki akses ke rumah sakit ini")
	ErrInvalidFormVersionIdHospitalId       = Error(http.StatusForbidden, "formulir ini tidak tersedia untuk rumah sakit ini")
	ErrInvalidFormEntryRoleId               = Error(http.StatusForbidden, "anda tidak dapat mengisi formulir ini")
	ErrInvalidFormEntryAssignment           = Error(http.StatusForbidden, "anda tidak ditugaskan untuk mengisi formulir ini")
	ErrDeleteSystemInitiatedNotification    = Error(http.StatusBadRequest, "tidak dapat menghapus notifikasi yang diinisiasi sistem")
	ErrInvalidVehicleNumber                 = Error(http.StatusBadRequest, "plat nomor kendaraan tidak valid")
	ErrExpiredFormApprovalCode              = Error(http.StatusUnauthorized, "kode persetujuan sudah kadaluarsa, silakan login untuk melanjutkan proses")
	ErrInvalidFormVersionId                 = Error(http.StatusBadRequest, "invalid form version ID")
	ErrInvalidReportType                    = Error(http.StatusInternalServerError, "invalid report type")
	ErrEmptyRemarksIncompleteForm           = Error(http.StatusBadRequest, "mohon isi catatan untuk penyerahan formulir yang tidak lengkap")
	ErrEmptyFormEntrySubmission             = Error(http.StatusBadRequest, "tidak dapat menyerahkan formulir kosong")
	ErrMandatoryEmptyFormEntry              = Error(http.StatusBadRequest, "terdapat isian wajib yang kosong")
	ErrIncompatibleFormSectionEntry         = Error(http.StatusBadRequest, "form section entry doesn't have the same number as form version's")
	ErrIncompatibleFormMatrixRowEntry       = Error(http.StatusBadRequest, "form matrix row entry doesn't have the same number as form version's")
	ErrIncompatibleFormMatrixColumnEntry    = Error(http.StatusBadRequest, "form matrix column entry doesn't have the same number as form version's")
	ErrInvalidEntrySection                  = Error(http.StatusBadRequest, "invalid entry section")
	ErrActiveEntryExists                    = Error(http.StatusBadRequest, "tidak dapat menyerahkan formulir kembali karena terdapat formulir yang masih dalam proses verifikasi")
	ErrEmptyRequiredParams                  = Error(http.StatusBadRequest, "required parameters are empty")
	ErrDoubleEntryValue                     = Error(http.StatusBadRequest, "entry value cannot have both single value and multiple values")
	ErrEmptyFormEntryApproval               = Error(http.StatusBadRequest, "tidak terdapat formulir yang memerlukan verifikasi anda saat ini")
	ErrInvalidRoleReporter                  = Error(http.StatusBadRequest, "peran yang dipilih tidak dapat mengisi formulir")
	ErrInvalidRoleVerificator               = Error(http.StatusBadRequest, "peran yang dipilih tidak dapat memverifikasi formulir")
	ErrDuplicateCronExpression              = Error(http.StatusBadRequest, "tidak dapat membuat notifikasi dengan lebih dari satu jadwal yang sama")
	ErrEmptyExcelImport                     = Error(http.StatusBadRequest, "tidak dapat menunggah dokumen kosong")
	ErrInvalidExcelImport                   = Error(http.StatusBadRequest, "dokumen yang diunggah tidak sesuai. Mohon gunakan templat yang sudah disediakan")
	ErrImportNotFound                       = Error(http.StatusNotFound, "data import tidak ditemukan")
	ErrFixImportEmptyDivision               = Error(http.StatusInternalServerError, "either division ID or new division should be filled")
	ErrFixImportEmptyPosition               = Error(http.StatusInternalServerError, "either position ID or new position should be filled")
	ErrEmptyValidImportEmployee             = Error(http.StatusBadRequest, "tidak ada data yang valid untuk sesi import karyawan ini")
	ErrInvalidWorkShiftCodeUse              = Error(http.StatusBadRequest, "tidak dapat menggunakan kode shift ini")
	ErrRestrictWorkShiftUpdateCode          = Error(http.StatusBadRequest, "tidak dapat mengubah kode shift ini")
	ErrRestrictWorkShiftDeletion            = Error(http.StatusBadRequest, "tidak dapat menghapus shift ini")
	ErrRestrictWorkShiftUpdateActivation    = Error(http.StatusBadRequest, "tidak dapat mengubah status aktivasi shift ini")
	ErrInvalidWorkShiftCode                 = Error(http.StatusBadRequest, "kode shift tidak valid")
	ErrConflictEmployeeSchedule             = Error(http.StatusBadRequest, "terdapat jadwal karyawan yang bentrok")
	ErrInsufficientGenerateScheduleEmployee = Error(http.StatusBadRequest, "tidak dapat membuat jadwal secara otomatis untuk jabatan yang memiliki kurang dari 4 karyawan")
)
