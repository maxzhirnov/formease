// src/lib/utils/showError.ts
import { errorStore } from '$lib/stores/errorStore';

export function showError(message: string) {
    errorStore.set(message);
}
