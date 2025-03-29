import type { Metadata } from 'next'
import localFont from 'next/font/local'
import './globals.css'
import { Nav } from '../components/nav/Nav'
import { AuthWrapper } from '@/components/AuthWrapper'

const geistSans = localFont({
  src: './fonts/GeistVF.woff',
  variable: '--font-geist-sans',
  weight: '100 900'
})
const geistMono = localFont({
  src: './fonts/GeistMonoVF.woff',
  variable: '--font-geist-mono',
  weight: '100 900'
})

export const metadata: Metadata = {
  title: 'Tsuki',
  description: 'Lightweight manga and comics media server'
}

export default function RootLayout ({
  children
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
		<meta name="wails-options" content="noautoinject" />
		<script src="/wails/ipc.js" async></script>
		<script src="/wails/runtime.js" async></script>
      </head>
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <AuthWrapper>
          <Nav />
          {children}
        </AuthWrapper>
      </body>
    </html>
  )
}
