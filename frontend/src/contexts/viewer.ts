import { Viewer } from "@/types/anilist"
import { createContext } from 'react'

export const ViewerContext = createContext<Viewer | null>(null)

export const ViewerProvider = ViewerContext.Provider
