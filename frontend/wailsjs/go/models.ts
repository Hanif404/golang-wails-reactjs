export namespace auth {
	
	export class Login {
	    email: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new Login(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.email = source["email"];
	        this.password = source["password"];
	    }
	}

}

export namespace student {
	
	export class Student {
	    id: number;
	    nis: string;
	    name: string;
	    gender: string;
	    phone: string;
	    email: string;
	    kelas: string;
	    angkatan: string;
	    alamat: string;
	    tempat_lahir: string;
	    tanggal_lahir: string;
	    nama_ayah: string;
	    nama_ibu: string;
	    status: string;
	    synced: boolean;
	    created_by: string;
	    created_at: number;
	
	    static createFrom(source: any = {}) {
	        return new Student(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nis = source["nis"];
	        this.name = source["name"];
	        this.gender = source["gender"];
	        this.phone = source["phone"];
	        this.email = source["email"];
	        this.kelas = source["kelas"];
	        this.angkatan = source["angkatan"];
	        this.alamat = source["alamat"];
	        this.tempat_lahir = source["tempat_lahir"];
	        this.tanggal_lahir = source["tanggal_lahir"];
	        this.nama_ayah = source["nama_ayah"];
	        this.nama_ibu = source["nama_ibu"];
	        this.status = source["status"];
	        this.synced = source["synced"];
	        this.created_by = source["created_by"];
	        this.created_at = source["created_at"];
	    }
	}

}

