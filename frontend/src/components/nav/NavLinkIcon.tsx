'use client'

import { IconProp } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { MouseEventHandler } from 'react'

interface Props {
    icon: IconProp
    href?: string | undefined
    onClick?: MouseEventHandler<HTMLDivElement> | undefined
}

export function NavLinkIcon ({ icon, href, onClick }: Props) {
  const pathname = usePathname()

  const classes = () => {
    return pathname === href ? 'text-foreground bg-foreground/10' : 'text-foreground/75 hover:text-foreground hover:bg-foreground/10'
  }

  const linkOrButton = () => {
    if (href === undefined) {
      return (
                <div onClick={onClick} className={`rounded px-4 py-2 cursor-pointer ${classes()} transition duration-300 ease-in-out`}>
                    <FontAwesomeIcon icon={icon} />
                </div>
      )
    }

    return (
            <Link className={`rounded px-4 py-2 ${classes()} transition duration-300 ease-in-out`} href={href ?? ''}>
                <FontAwesomeIcon icon={icon} />
            </Link>
    )
  }

  return linkOrButton()
}
