'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'

interface Props {
    text: string
    href: string
}

export function NavLink ({ text, href }: Props) {
  const pathname = usePathname()

  const classes = () => {
    console.log(pathname)
    return pathname.replaceAll('/', '') === href.replaceAll('/', '') ? 'text-foreground bg-foreground/10' : 'text-foreground/75 hover:text-foreground hover:bg-foreground/10'
  }

  return (
        <Link className={`rounded px-4 py-2 cursor-pointer h-fit ${classes()} transition duration-300 ease-in-out`} href={href}>{text}</Link>
  )
}
