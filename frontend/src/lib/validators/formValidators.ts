export type ValidatorName = 'email' | 'zipCode' | 'phone' // Add more as needed

export const validators: Record<ValidatorName, RegExp> = {
  email: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
  zipCode: /^\d{5}(-\d{4})?$/,
  phone: /^(\+\d{1,2}\s?)?1?\-?\.?\s?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$/
  // Add more validators as needed
};

export function validate(value: string, validatorName: ValidatorName): boolean {
  const validator = validators[validatorName];
  return validator ? validator.test(value) : false;
}
