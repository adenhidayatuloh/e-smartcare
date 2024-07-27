package siswarepository

import (
	"esmartcare/entity"
	"esmartcare/pkg/errs"
)

type SiswaRepository interface {
	CreateSiswa(siswa *entity.Siswa) (*entity.Siswa, errs.MessageErr)
	UpdateSiswa(oldSiswa *entity.Siswa, newSiswa *entity.Siswa) (*entity.Siswa, errs.MessageErr)
	//UpdatePhoto(url string) (*entity.Siswa, errs.MessageErr)
	// GetAllTasks() ([]entity.Task, errs.MessageErr)
	// GetTaskByID(id uint) (*entity.Task, errs.MessageErr)
	// UpdateTask(oldTask *entity.Task, newTask *entity.Task) (*entity.Task, errs.MessageErr)
	// UpdateTaskStatus(id uint, newStatus bool) (*entity.Task, errs.MessageErr)
	// UpdateTaskCategory(id uint, newCategoryID uint) (*entity.Task, errs.MessageErr)
	// DeleteTask(id uint) errs.MessageErr

	GetSiswaByEmail(email string) (*entity.Siswa, errs.MessageErr)
	GetAllSiswaWithPemeriksaan() ([]entity.Siswa_pemeriksaan, errs.MessageErr)
}
