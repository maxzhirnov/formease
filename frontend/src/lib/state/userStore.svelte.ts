// import { browser } from '$app/environment';

// interface UserState {
//     isAuthenticated: boolean;
//     id: string;
//     email: string;
// }

// // Default state
// const defaultState: UserState = {
//     isAuthenticated: false,
//     id: '',
//     email: '',
// };

// // Initialize state from localStorage only in browser
// const initialState = browser 
//     ? JSON.parse(localStorage.getItem('user') || JSON.stringify(defaultState))
//     : defaultState;

// export const userState = $state(initialState);

// // Only sync to localStorage in browser environment
// if (browser) {
//     $effect(() => {
//         localStorage.setItem('user', JSON.stringify(userState));
//     });
// }
