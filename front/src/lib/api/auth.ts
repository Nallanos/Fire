import { API_URL } from "./index";

type authParams = {
    name: string;
    password: string;
    email: string
}

type tokenParams = {
    email: string
}

export async function Login(authParams: authParams): Promise<{ data?: string, success: boolean; message?: string }> {
    const req = new Request(API_URL + `/signIn`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(authParams)
    });
    try {
        const res = await fetch(req);
        if (res.ok) {
            const data = await res.json();
            return {
                data: data,
                success: true,
                message: "Login successful"
            };
        } else {
            const errorData = await res.json();
            return {
                data: errorData,
                success: false,
                message: errorData.message || "Login failed"
            };
        }
    } catch (err) {
        console.error(err);
        return {
            success: false,
            message: "An error occurred while trying to log in"
        };
    }
}

export async function GetToken(tokenParams: tokenParams) {
    const req = new Request(API_URL + `/signIn`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(tokenParams)
    });
    try {
        const res = await fetch(req);
        console.log(res.status);
        const data = await res.json();
        console.log(data);
        return data
    } catch (err) {
        console.error(err);
    }
}
export async function SignUp(authParams: authParams): Promise<{ success: boolean, message?: string }> {
    const req = new Request(API_URL + `/signUp`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(authParams)
    });
    try {
        const res = await fetch(req);
        if (res.ok) {
            return {
                success: true,
                message: "sign up successfuly"
            }
        } else {
            const Errordata = await res.json();
            return {
                success: true,
                message: "sign up successfuly"
            }
        }
    } catch (err) {
        console.error(err);
        return {
            success: false,
            message: "An error occurred while trying to sign up"
        }
    }
}
