package gen

import "github.com/lithammer/shortuuid/v4"

func GenShortUuid() string {
	return shortuuid.New()
}

func GenBorrowerId() string {
	return "bor_" + GenShortUuid()
}

func GenBookId() string {
	return "book_" + GenShortUuid()
}

func GenBookTitleId() string {
	return "btl_" + GenShortUuid()
}

func GenLibrarianId() string {
	return "lrn_" + GenShortUuid()
}

func GenAuthorId() string {
	return "athr_" + GenShortUuid()
}
