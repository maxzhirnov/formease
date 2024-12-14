export type ThemeName = 
  | 'light' 
  | 'dark' 
  | 'ocean' 
  | 'sunset' 
  | 'forest'
  | 'tech'
  | 'cream'
  | 'deepBlue'
  | 'mint'
  | 'purple'
  | 'pinkTeal'
  | 'sunsetOrange'
  | 'deepPurple'
  | 'vibrantYellow'
  | 'neonGreen';

export interface Theme {
  backgroundColor: string;
  fontColor: string;
}

export const predefinedThemes: Record<ThemeName, Theme> = {
    light: { backgroundColor: "#FFFFFF", fontColor: "#000000" },
    dark: { backgroundColor: "#2D0A31", fontColor: "#FFFFFF" },
    ocean: { backgroundColor: "#2193b0", fontColor: "#FFFFFF" },
    sunset: { backgroundColor: "#FF5841", fontColor: "#FFFFFF" },
    forest: { backgroundColor: "#228B22", fontColor: "#FFFFFF" },
    tech: { backgroundColor: "#1D1842", fontColor: "#FFFFFF" },
    cream: { backgroundColor: "#F2F0EA", fontColor: "#333333" },
    deepBlue: { backgroundColor: "#000428", fontColor: "#FFFFFF" },
    mint: { backgroundColor: "#00b09b", fontColor: "#FFFFFF" },
    purple: { backgroundColor: "#7C3AED", fontColor: "#FFFFFF" },
    pinkTeal: { backgroundColor: "#FF78AC", fontColor: "#FFFFFF" },
    sunsetOrange: { backgroundColor: "#FFAB00", fontColor: "#333333" },
    deepPurple: { backgroundColor: "#8E0D3C", fontColor: "#FFFFFF" },
    vibrantYellow: { backgroundColor: "#FFD43A", fontColor: "#582C12" },
    neonGreen: { backgroundColor: "#00D494", fontColor: "#FFFFFF" }
};
