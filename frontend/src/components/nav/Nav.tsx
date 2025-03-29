'use client'

import { faPuzzlePiece, faRightFromBracket } from '@fortawesome/free-solid-svg-icons'
import { NavLink } from './NavLink'
import { Logo } from '../svg/logo'
import { NavLinkIcon } from './NavLinkIcon'

export function Nav () {
  const logout = () => {
    localStorage.setItem('tsuki_account', '')
    window.location.href = '/'
  }

  return (
    <nav className="bg-background/50 backdrop-blur-sm flex justify-between px-8 py-6 fixed top-0 w-full border-b-2 border-foreground z-10">
        <div className="flex gap-2 items-center">
            <Logo width={50} height={50} className="mr-4" />
            <NavLink text="Dashboard" href="/" />
            <NavLink text="Manga" href="/manga" />
            <NavLink text="Comics" href="/comics" />
        </div>

        <div className="flex gap-2 items-center">
            <NavLinkIcon icon={faPuzzlePiece} href="/repositories" />
            <NavLinkIcon icon={faRightFromBracket} onClick={logout} />
        </div>
    </nav>
  )
}
