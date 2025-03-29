'use client'

import { FullscreenCenter } from "./FullscreenCenter"

export function LoadingScreen () {
  return (
        <FullscreenCenter>
            <div className="z-[-10] absolute left-0 right-0 top-0 bottom-0 my-auto mx-auto w-[25px] h-[25px] loader border-[2px] border-t-[5px] border-foreground rounded-full animate-spin"></div>
        </FullscreenCenter>
  )
}
