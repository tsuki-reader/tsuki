export namespace backend {
	
	export class AnilistStatus {
	    authenticated: boolean;
	    viewer?: types.ALViewer;
	    client_id: string;
	
	    static createFrom(source: any = {}) {
	        return new AnilistStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.authenticated = source["authenticated"];
	        this.viewer = this.convertValues(source["viewer"], types.ALViewer);
	        this.client_id = source["client_id"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class LoginStatus {
	    success: boolean;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new LoginStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	    }
	}

}

export namespace models {
	
	export class Account {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    anilist_token: string;
	    anilist_name: string;
	    username: string;
	
	    static createFrom(source: any = {}) {
	        return new Account(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.anilist_token = source["anilist_token"];
	        this.anilist_name = source["anilist_name"];
	        this.username = source["username"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace types {
	
	export class ALAvatar {
	    large: string;
	    medium: string;
	
	    static createFrom(source: any = {}) {
	        return new ALAvatar(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.large = source["large"];
	        this.medium = source["medium"];
	    }
	}
	export class ALCoverImage {
	    extraLarge: string;
	    large: string;
	    medium: string;
	    color: string;
	
	    static createFrom(source: any = {}) {
	        return new ALCoverImage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.extraLarge = source["extraLarge"];
	        this.large = source["large"];
	        this.medium = source["medium"];
	        this.color = source["color"];
	    }
	}
	export class ALDate {
	    year: number;
	    month: number;
	    day: number;
	
	    static createFrom(source: any = {}) {
	        return new ALDate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.year = source["year"];
	        this.month = source["month"];
	        this.day = source["day"];
	    }
	}
	export class ALTitle {
	    romaji: string;
	    english: string;
	    native: string;
	
	    static createFrom(source: any = {}) {
	        return new ALTitle(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.romaji = source["romaji"];
	        this.english = source["english"];
	        this.native = source["native"];
	    }
	}
	export class ALManga {
	    id: number;
	    title: ALTitle;
	    startDate: ALDate;
	    status: string;
	    chapters: number;
	    volumes: number;
	    coverImage: ALCoverImage;
	    bannerImage: string;
	    description: string;
	    genres: string[];
	
	    static createFrom(source: any = {}) {
	        return new ALManga(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = this.convertValues(source["title"], ALTitle);
	        this.startDate = this.convertValues(source["startDate"], ALDate);
	        this.status = source["status"];
	        this.chapters = source["chapters"];
	        this.volumes = source["volumes"];
	        this.coverImage = this.convertValues(source["coverImage"], ALCoverImage);
	        this.bannerImage = source["bannerImage"];
	        this.description = source["description"];
	        this.genres = source["genres"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ALMediaList {
	    progress: number;
	    completedAt: ALDate;
	    startedAt: ALDate;
	    notes: string;
	    score: number;
	    status: string;
	    media: ALManga;
	
	    static createFrom(source: any = {}) {
	        return new ALMediaList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.progress = source["progress"];
	        this.completedAt = this.convertValues(source["completedAt"], ALDate);
	        this.startedAt = this.convertValues(source["startedAt"], ALDate);
	        this.notes = source["notes"];
	        this.score = source["score"];
	        this.status = source["status"];
	        this.media = this.convertValues(source["media"], ALManga);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ALMediaListGroup {
	    name: string;
	    isCustomList: boolean;
	    isSplitCustomList: boolean;
	    status: string;
	    entries: ALMediaList[];
	
	    static createFrom(source: any = {}) {
	        return new ALMediaListGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.isCustomList = source["isCustomList"];
	        this.isSplitCustomList = source["isSplitCustomList"];
	        this.status = source["status"];
	        this.entries = this.convertValues(source["entries"], ALMediaList);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class ALViewer {
	    name: string;
	    bannerImage: string;
	    avatar: ALAvatar;
	
	    static createFrom(source: any = {}) {
	        return new ALViewer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.bannerImage = source["bannerImage"];
	        this.avatar = this.convertValues(source["avatar"], ALAvatar);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

