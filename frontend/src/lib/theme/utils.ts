export function getContrastColor(backgroundColor: string): string {
    // Convert hex to RGB
    const hex = backgroundColor.replace('#', '');
    const r = parseInt(hex.substr(0, 2), 16);
    const g = parseInt(hex.substr(2, 2), 16);
    const b = parseInt(hex.substr(4, 2), 16);
  
    // Calculate luminance
    const luminance = (0.299 * r + 0.587 * g + 0.114 * b) / 255;
  
    // Return contrasting color
    return luminance > 0.5 ? 'blue-500' : 'blue-200';
}

export const randomBrightColor = () => {
    const brightColors = [
      'bg-red-500', 
      'bg-blue-500', 
      'bg-green-500', 
      'bg-purple-500', 
      'bg-pink-500', 
      'bg-indigo-500', 
      'bg-yellow-500', 
      'bg-orange-500'
    ];
    return brightColors[Math.floor(Math.random() * brightColors.length)];
  };