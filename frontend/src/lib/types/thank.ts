export interface ThankYouMessage {
    title: string;
    subtitle: string;
    icon: string;
    button: Button;
}
  
export interface Button {
  text: string;
  url: string;
  newTab: boolean;
}

export const defaultThankYouMessage: ThankYouMessage = {
  title: 'Спасибо!',
  subtitle: 'за то что прошли наш опрос!',
  icon: '🥳',
  button: {
      text: 'На наш сайт',
      url: 'https://ya.ru',
      newTab: true
  }
}
