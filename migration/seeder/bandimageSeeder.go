package seeder

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khrongpop/BandSquareAPI/model"
)

func BandImageSeed(db *gorm.DB) {
	fmt.Println("seed BandImage ...")
	db.Create(&model.BandImage{
		BandtypeID: 1,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/a4b6d58c-23d3-4f66-bc49-5e9d84cf4c1b.png`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/a4b6d58c-23d3-4f66-bc49-5e9d84cf4c1b.png`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 1,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/0a13aaec-7b73-4905-9e85-4665aec5505b.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/0a13aaec-7b73-4905-9e85-4665aec5505b.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 1,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/914000f3-6378-485d-b312-818f0e55d7b0.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/914000f3-6378-485d-b312-818f0e55d7b0.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 2,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/b17f7f5b-ce46-4a69-afa0-232c0818f48f.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/b17f7f5b-ce46-4a69-afa0-232c0818f48f.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 2,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/eebd5027-10fd-4975-8cf0-054ac43bacd8.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/eebd5027-10fd-4975-8cf0-054ac43bacd8.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 2,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/f13f2a07-b84c-433b-af20-ee7cb7eb1432.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/f13f2a07-b84c-433b-af20-ee7cb7eb1432.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 3,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/df227567-0d17-4452-8802-d994adf2ae16.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/df227567-0d17-4452-8802-d994adf2ae16.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 3,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7ef6a8d9-a51f-443a-bb9b-7a7d13b18049.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7ef6a8d9-a51f-443a-bb9b-7a7d13b18049.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 3,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/92ae6490-d076-4016-a117-5f3a72d44315.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/92ae6490-d076-4016-a117-5f3a72d44315.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 4,
		Image:      `https://lensod-user-statics.s3.amazonaws.com/backend_1539071035557.png`,
		Thumbnail:  `https://lensod-user-statics.s3.amazonaws.com/backend_1539071035557.png`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 4,
		Image:      `https://lensod-user-statics.s3.amazonaws.com/backend_1539071247956.png`,
		Thumbnail:  `https://lensod-user-statics.s3.amazonaws.com/backend_1539071247956.png`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 4,
		Image:      `https://lensod-user-statics.s3.amazonaws.com/backend_1539071231421.png`,
		Thumbnail:  `https://lensod-user-statics.s3.amazonaws.com/backend_1539071231421.png`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 5,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/6dc25999-e355-4ee6-8dc0-d17cb401c5e9.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/6dc25999-e355-4ee6-8dc0-d17cb401c5e9.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 5,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/ff75e843-822b-4453-8126-1921ac0b8ce8.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/ff75e843-822b-4453-8126-1921ac0b8ce8.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 5,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7a5ee10b-e515-432d-b7c1-ddce4f6fbaee.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7a5ee10b-e515-432d-b7c1-ddce4f6fbaee.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 6,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/93dafa11-c12c-4a58-908b-47cad072208d.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/93dafa11-c12c-4a58-908b-47cad072208d.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 6,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7cad4d0c-eaf4-47c3-ac1c-b186e38e25a0.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7cad4d0c-eaf4-47c3-ac1c-b186e38e25a0.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 7,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/13a68f06-8b5e-4b08-8ed6-d4afad7da7d9.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/13a68f06-8b5e-4b08-8ed6-d4afad7da7d9.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 7,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/efd4a23a-ce72-490f-a315-16dc49f61189.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/efd4a23a-ce72-490f-a315-16dc49f61189.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 7,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/f34d23e7-9800-4a39-a3ba-951ba07fa555.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/f34d23e7-9800-4a39-a3ba-951ba07fa555.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 8,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/96956a33-98b3-41eb-9ec3-353f4bdb61b0.png`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/96956a33-98b3-41eb-9ec3-353f4bdb61b0.png`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 8,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/2d54e352-636e-4355-aa54-071707399fec.png`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/2d54e352-636e-4355-aa54-071707399fec.png`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 8,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/018e21c3-4c05-4322-b0c2-4267fadbe5d7.png`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/018e21c3-4c05-4322-b0c2-4267fadbe5d7.png`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 9,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/20b86b43-1e88-41ed-b63f-43c2fba5abf3.jpe`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/20b86b43-1e88-41ed-b63f-43c2fba5abf3.jpe`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 9,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/848a34c9-de3d-4fe9-aa1c-b9089fbbcffc.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/848a34c9-de3d-4fe9-aa1c-b9089fbbcffc.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 9,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/696ce866-0240-4c48-9630-fc976db2c6f9.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/696ce866-0240-4c48-9630-fc976db2c6f9.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 10,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/86d8f6d4-7ec8-4633-89ba-027dfd4404ab.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/86d8f6d4-7ec8-4633-89ba-027dfd4404ab.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 10,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/00bda6d5-0f5e-4913-8ff3-0016ad720b8a.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/00bda6d5-0f5e-4913-8ff3-0016ad720b8a.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 10,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/8984d8a8-c370-4763-9c4a-d471e7dbcb04.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/8984d8a8-c370-4763-9c4a-d471e7dbcb04.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 11,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/b1809ded-0eb2-48ae-a363-b0463c05cf68.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/b1809ded-0eb2-48ae-a363-b0463c05cf68.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 11,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/fe81d74e-488f-4d5f-a35f-790ce82e7744.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/fe81d74e-488f-4d5f-a35f-790ce82e7744.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 11,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/e211dd5d-25ba-4fbc-a854-65c56885857f.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/e211dd5d-25ba-4fbc-a854-65c56885857f.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 12,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/36cb8f8d-245f-40af-80b1-753eb8a23ba0.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/36cb8f8d-245f-40af-80b1-753eb8a23ba0.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 12,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/1de78d25-fff7-49ac-93e7-51a085be63f8.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/1de78d25-fff7-49ac-93e7-51a085be63f8.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 12,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/dc5f54a5-3936-4a20-b630-dece0fa433c0.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/dc5f54a5-3936-4a20-b630-dece0fa433c0.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 13,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/411cd384-3f54-497f-bded-2378cd872735.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/411cd384-3f54-497f-bded-2378cd872735.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 13,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/90c943dc-d5f5-41f9-9b57-d542730d3db8.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/90c943dc-d5f5-41f9-9b57-d542730d3db8.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 13,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/28758fc6-5e06-4366-bc5b-11f2b125d7fe.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/28758fc6-5e06-4366-bc5b-11f2b125d7fe.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 14,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/6fb77f91-b06b-45fc-ba9b-05b369adb0e2.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/6fb77f91-b06b-45fc-ba9b-05b369adb0e2.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 14,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/4fb304aa-9b93-4032-a16b-6cfd0b6de864.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/4fb304aa-9b93-4032-a16b-6cfd0b6de864.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 14,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/cd4b3776-7821-4976-abfe-27a519d613d4.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/cd4b3776-7821-4976-abfe-27a519d613d4.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 15,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/4f4804ef-d294-4708-8fcd-62f03fef3a48.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/4f4804ef-d294-4708-8fcd-62f03fef3a48.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 15,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7994ace6-e118-4339-9e0a-a751a450636c.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7994ace6-e118-4339-9e0a-a751a450636c.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 15,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/b0df4efe-cf3c-49bd-a7fd-e7732bb21694.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/b0df4efe-cf3c-49bd-a7fd-e7732bb21694.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 16,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/48d24064-6501-4bd0-af94-899ecae79535.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/48d24064-6501-4bd0-af94-899ecae79535.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 16,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7bff3bf2-a5bb-4b68-aeec-268b9cbd19f6.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/7bff3bf2-a5bb-4b68-aeec-268b9cbd19f6.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 16,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/bcea1b83-952d-4df2-9dcc-e2a84152e73c.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/bcea1b83-952d-4df2-9dcc-e2a84152e73c.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 17,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/51f95396-dbd5-492d-825e-f9ede097b6d3.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/51f95396-dbd5-492d-825e-f9ede097b6d3.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 17,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/ce19eee4-95fa-48a4-875f-765d518a4560.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/ce19eee4-95fa-48a4-875f-765d518a4560.jpeg`,
	})
	db.Create(&model.BandImage{
		BandtypeID: 17,
		Image:      `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/0be1ed62-539c-4c32-bd6c-ab5363b1436f.jpeg`,
		Thumbnail:  `https://lensod-user-statics.s3-ap-southeast-1.amazonaws.com/0be1ed62-539c-4c32-bd6c-ab5363b1436f.jpeg`,
	})

}
