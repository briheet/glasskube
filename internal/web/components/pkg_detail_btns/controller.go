package pkg_detail_btns

import (
	"fmt"

	"github.com/glasskube/glasskube/api/v1alpha1"
	"github.com/glasskube/glasskube/pkg/client"
)

const TemplateId = "pkg-detail-btns"

type pkgDetailBtnsInput struct {
	ContainerId     string
	PackageName     string
	Status          *client.PackageStatus
	Manifest        *v1alpha1.PackageManifest
	UpdateAvailable bool
	Pkg             *v1alpha1.Package
}

func getId(pkgName string) string {
	return fmt.Sprintf("%v-%v", TemplateId, pkgName)
}

func ForPkgDetailBtns(
	pkgName string,
	status *client.PackageStatus,
	manifest *v1alpha1.PackageManifest,
	pkg *v1alpha1.Package,
	updateAvailable bool,
) *pkgDetailBtnsInput {
	id := getId(pkgName)
	return &pkgDetailBtnsInput{
		ContainerId:     id,
		PackageName:     pkgName,
		Status:          status,
		Manifest:        manifest,
		UpdateAvailable: updateAvailable,
		Pkg:             pkg,
	}
}
