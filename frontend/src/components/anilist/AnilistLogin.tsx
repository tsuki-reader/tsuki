'use client'

import { FullscreenCenter } from '../FullscreenCenter'
import { AnilistLogo } from '../svg/AnilistLogo'
import { Logo } from '../svg/Logo'

interface Props {
  clientId: string
}

export function AnilistLogin ({ clientId }: Props) {
  return (
    <FullscreenCenter>
      <div className="flex flex-col items-center text-center rounded-lg border-2 border-foreground m-4 p-16">
        <Logo width={200} height={150} className="ml-4" />
        <h1 className="text-2xl font-bold p-8">Login with Anilist</h1>
        <a
          className="bg-[#3db4f2] flex gap-2 whitespace-nowrap items-center rounded px-4 py-1 font-bold hover:brightness-90"
          href={`https://anilist.co/api/v2/oauth/authorize?client_id=${clientId}&response_type=token`}
        >
          <AnilistLogo width={30} height={30} />
          Login
        </a>
      </div>
    </FullscreenCenter>
  )
}
