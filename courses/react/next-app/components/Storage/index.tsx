import { useState, useEffect } from "react";

const userKey = 'user';

const Storage = {
    saveUser(data: { email: string, password: string }): void {
        window.localStorage.setItem(userKey, JSON.stringify(data));
    },

    hasUser(): boolean {
        const user: string | null = window.localStorage.getItem(userKey);
        return user !== null;
    },

    available(): boolean {
        return localStorage !== undefined;
    }
};

export default Storage;

export function useLogged() {
    const [userAvailable, setUserAvailable] = useState(false);
    useEffect(() => {
        if (Storage.available()) {
            setUserAvailable(Storage.hasUser());
        }
    });
    return userAvailable;
};
