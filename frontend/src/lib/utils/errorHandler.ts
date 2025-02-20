import { errorStore } from '$lib/state/errorStore';

// Declare overload signatures
export function showError(message: string): void;
export function showError(err: Error): void;
export function showError(err: unknown, defaultMessage?: string): void;

// Implement the function
export function showError(input: string | Error | unknown, defaultMessage: string = 'An unexpected error occurred'): void {
    let message: string;
    
    if (typeof input === 'string') {
        message = input;
    } else if (input instanceof Error) {
        message = input.message;
    } else {
        message = defaultMessage;
    }
    
    errorStore.set(message);
}
