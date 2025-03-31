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
	export class MangaShowResponse {
	    media_list?: types.ALMediaList;
	    chapters: models.Chapter[];
	
	    static createFrom(source: any = {}) {
	        return new MangaShowResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.media_list = this.convertValues(source["media_list"], types.ALMediaList);
	        this.chapters = this.convertValues(source["chapters"], models.Chapter);
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
	export class MappingAssignResponse {
	    media_list: types.ALMediaList;
	    chapters: models.Chapter[];
	
	    static createFrom(source: any = {}) {
	        return new MappingAssignResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.media_list = this.convertValues(source["media_list"], types.ALMediaList);
	        this.chapters = this.convertValues(source["chapters"], models.Chapter);
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
	export class MappingChapterPagesResponse {
	    pages: providers.Page[];
	    installed_provider: models.InstalledProvider;
	
	    static createFrom(source: any = {}) {
	        return new MappingChapterPagesResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pages = this.convertValues(source["pages"], providers.Page);
	        this.installed_provider = this.convertValues(source["installed_provider"], models.InstalledProvider);
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

export namespace extensions {
	
	export class Provider {
	    name: string;
	    id: string;
	    file: string;
	    icon: string;
	    installed: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Provider(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.id = source["id"];
	        this.file = source["file"];
	        this.icon = source["icon"];
	        this.installed = source["installed"];
	    }
	}
	export class Repository {
	    name: string;
	    id: string;
	    logo: string;
	    url: string;
	    manga_providers: Provider[];
	    comic_providers: Provider[];
	
	    static createFrom(source: any = {}) {
	        return new Repository(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.id = source["id"];
	        this.logo = source["logo"];
	        this.url = source["url"];
	        this.manga_providers = this.convertValues(source["manga_providers"], Provider);
	        this.comic_providers = this.convertValues(source["comic_providers"], Provider);
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
	export class InstalledProvider {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    name: string;
	    providerId: string;
	    repositoryId: string;
	    providerType: string;
	
	    static createFrom(source: any = {}) {
	        return new InstalledProvider(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.name = source["name"];
	        this.providerId = source["providerId"];
	        this.repositoryId = source["repositoryId"];
	        this.providerType = source["providerType"];
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
	export class Chapter {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    title: string;
	    external_id: string;
	    provider: string;
	    absolute_number: number;
	    installed_provider: InstalledProvider;
	
	    static createFrom(source: any = {}) {
	        return new Chapter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.title = source["title"];
	        this.external_id = source["external_id"];
	        this.provider = source["provider"];
	        this.absolute_number = source["absolute_number"];
	        this.installed_provider = this.convertValues(source["installed_provider"], InstalledProvider);
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
	
	export class Mapping {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    anilistId: number;
	    externalId: string;
	    progress: number;
	    chapters: number;
	    installedProvider: InstalledProvider;
	
	    static createFrom(source: any = {}) {
	        return new Mapping(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.anilistId = source["anilistId"];
	        this.externalId = source["externalId"];
	        this.progress = source["progress"];
	        this.chapters = source["chapters"];
	        this.installedProvider = this.convertValues(source["installedProvider"], InstalledProvider);
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

export namespace providers {
	
	export class Page {
	    provider: string;
	    image_url: string;
	    page_number: number;
	
	    static createFrom(source: any = {}) {
	        return new Page(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.provider = source["provider"];
	        this.image_url = source["image_url"];
	        this.page_number = source["page_number"];
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
	    mapping?: models.Mapping;
	
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
	        this.mapping = this.convertValues(source["mapping"], models.Mapping);
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

