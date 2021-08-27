const userKey = 'user';

const Storage = {
    saveUser(data: { email: string, password: string }): void {
        localStorage.setItem(userKey, JSON.stringify(data));
    },

    hasUser(): boolean {
        const user: string | null = localStorage.getItem(userKey);
        return user !== null;
    },
};

export default Storage;
