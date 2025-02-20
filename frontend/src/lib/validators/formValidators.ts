export type ValidatorName = 
  | 'email' 
  | 'phone'
  | 'name'
  | 'username'
  | 'password'
  | 'url'
  | 'date'
  | 'time'
  | 'datetime'
  | 'color'
  | 'ipv4'
  | 'ipv6'
  | 'mac'
  | 'latitude'
  | 'longitude'
  | 'number'
  | 'integer'
  | 'float'
  | 'alphanumeric'
  | 'text'
  | 'cyrillicText'
  | 'passport'
  | 'inn'
  | 'snils'
  | 'creditCard';

export const validators: Record<ValidatorName, RegExp> = {
  // Basic
  email: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
  phone: /^(\+7|7|8)?[\s\-]?\(?[489][0-9]{2}\)?[\s\-]?[0-9]{3}[\s\-]?[0-9]{2}[\s\-]?[0-9]{2}$/,
  
  // Names and usernames
  name: /^[A-Za-zА-Яа-яЁё\s'-]{2,50}$/,
  username: /^[a-zA-Z0-9_-]{3,20}$/,
  
  // Password (min 8 chars, at least one number, one uppercase, one lowercase)
  password: /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$/,
  
  // URLs and web
  url: /^(https?:\/\/)?([\da-z.-]+)\.([a-z.]{2,6})([/\w .-]*)*\/?$/,
  
  // Dates and times
  date: /^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$/,
  time: /^([01]\d|2[0-3]):([0-5]\d)$/,
  datetime: /^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])T([01]\d|2[0-3]):([0-5]\d)$/,
  
  // Colors
  color: /^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$/,
  
  // Network
  ipv4: /^(\d{1,3}\.){3}\d{1,3}$/,
  ipv6: /^([0-9A-Fa-f]{1,4}:){7}[0-9A-Fa-f]{1,4}$/,
  mac: /^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$/,
  
  // Coordinates
  latitude: /^-?([1-8]?\d(?:\.\d{1,})?|90(?:\.0{1,})?)$/,
  longitude: /^-?((?:1[0-7]|[1-9])?\d(?:\.\d{1,})?|180(?:\.0{1,})?)$/,
  
  // Numbers
  number: /^-?\d*\.?\d+$/,
  integer: /^-?\d+$/,
  float: /^-?\d*\.\d+$/,
  
  // Text
  alphanumeric: /^[a-zA-Z0-9]+$/,
  text: /^.+$/,
  cyrillicText: /^[а-яА-ЯёЁ\s.,!?-]{1,}$/,
  
  // Russian documents
  passport: /^(\d{4})\s*(\d{6})$/,
  inn: /^\d{10}|\d{12}$/,
  snils: /^\d{3}-\d{3}-\d{3}\s\d{2}$/,
  
  // Payment
  creditCard: /^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|6(?:011|5[0-9]{2})[0-9]{12}|(?:2131|1800|35\d{3})\d{11})$/
};

// Helper function to test validation
export function validate(value: string, type: ValidatorName): boolean {
  console.log('validate', value, type);
  const validator = validators[type];
  if (!validator) {
    throw new Error(`Unknown validator type: ${type}`);
  }
  return validator.test(value);
}

// Optional: Add custom error messages
export const validationMessages: Record<ValidatorName, string> = {
  email: 'Please enter a valid email address',
  phone: 'Please enter a valid Russian phone number',
  name: 'Name should be 2-50 characters long and contain only letters, spaces, hyphens and apostrophes',
  username: 'Username should be 3-20 characters long and contain only letters, numbers, underscores and hyphens',
  password: 'Password must be at least 8 characters long and contain at least one number, one uppercase and one lowercase letter',
  url: 'Please enter a valid URL',
  date: 'Please enter a valid date in YYYY-MM-DD format',
  time: 'Please enter a valid time in HH:MM format',
  datetime: 'Please enter a valid datetime in YYYY-MM-DDTHH:MM format',
  color: 'Please enter a valid hex color code',
  ipv4: 'Please enter a valid IPv4 address',
  ipv6: 'Please enter a valid IPv6 address',
  mac: 'Please enter a valid MAC address',
  latitude: 'Please enter a valid latitude (-90 to 90)',
  longitude: 'Please enter a valid longitude (-180 to 180)',
  number: 'Please enter a valid number',
  integer: 'Please enter a valid integer',
  float: 'Please enter a valid decimal number',
  alphanumeric: 'Please use only letters and numbers',
  text: 'Please enter valid text',
  cyrillicText: 'Please enter text in Cyrillic',
  passport: 'Please enter a valid Russian passport number (4 digits space 6 digits)',
  inn: 'Please enter a valid INN (10 or 12 digits)',
  snils: 'Please enter a valid SNILS (XXX-XXX-XXX YY)',
  creditCard: 'Please enter a valid credit card number'
};
