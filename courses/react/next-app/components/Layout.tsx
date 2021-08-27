import React, { ReactNode } from 'react';

import Link from 'next/link';
import Head from 'next/head';

import { useLogged } from './Storage';

export interface LayoutProps {
    children?: ReactNode,
    title?: string,
};

export default function Layout({ children, title }: LayoutProps) {
    const logged = useLogged();

    const renderLinks = () => {
        return (
            <nav className="hidden md:block text-2xl space-x-4 md:mr-8">
                <span><Link href="/">Home</Link></span>
                { logged ? null : <span><Link href="/login">Login</Link></span> }
                { logged ? <span><Link href="/forecast">Forecast</Link></span> : null }
            </nav>
        );
    }

    return (
        <>
            <Head>
                <title>{ title }</title>
                <meta charSet="utf-8" />
                <meta name="viewport" content="initial-scale=1.0, width=device-width" />
            </Head>
            <header className="w-full top-0 fixed flex justify-center md:justify-between items-center h-24 bg-yellow-500 text-black">
                <span className="logo md:ml-8">Landing</span>
                { renderLinks() }
            </header>

            <main className="mt-24">
                { children }
            </main>
        </>
    );
};
