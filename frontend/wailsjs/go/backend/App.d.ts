// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {backend} from '../models';
import {context} from '../models';
import {types} from '../models';
import {extensions} from '../models';
import {models} from '../models';

export function AnilistLogin(arg1:string):Promise<backend.LoginStatus>;

export function AnilistStatus():Promise<backend.AnilistStatus>;

export function DomReady(arg1:context.Context):Promise<void>;

export function MangaIndex():Promise<Array<types.ALMediaListGroup>>;

export function MangaShow(arg1:number):Promise<backend.MangaShowResponse>;

export function ProvidersCreateOrUpdate(arg1:string,arg2:string,arg3:string):Promise<Array<extensions.Provider>>;

export function ProvidersDestroy(arg1:string,arg2:string,arg3:string):Promise<Array<extensions.Provider>>;

export function ProvidersIndex(arg1:string,arg2:string):Promise<Array<models.InstalledProvider>>;

export function RepositoriesCreate(arg1:string):Promise<Array<extensions.Repository>>;

export function RepositoriesDestroy(arg1:string):Promise<Array<extensions.Repository>>;

export function RepositoriesIndex():Promise<Array<extensions.Repository>>;

export function RepositoriesUpdate(arg1:string):Promise<extensions.Repository>;

export function SignIn(arg1:string,arg2:string):Promise<models.Account>;

export function SignOut():Promise<void>;
