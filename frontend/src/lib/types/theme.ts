export const themeOptions = {
    light: 'Light',
    dark: 'Dark',
    black: 'Black',
    white: 'White',
    ocean: 'Ocean',
    sunset: 'Sunset',
    forest: 'Forest',
    tech: 'Tech',
    cream: 'Cream',
    deepBlue: 'Deep Blue',
    mint: 'Mint',
    purple: 'Purple',
    pinkTeal: 'Pink Teal',
    sunsetOrange: 'Sunset Orange',
    deepPurple: 'Deep Purple',
    vibrantYellow: 'Vibrant Yellow',
    neonGreen: 'Neon Green'
  } as const
  
  export type ThemeName = keyof typeof themeOptions;
  
  export const themeOptionsList = Object.entries(themeOptions).map(([value, name]) => ({
    value: value as ThemeName,
    name: name as string
  }));
  
  export interface Theme {
    backgroundColor: string;
    fontColor: string;
    accentColor: string;
  }
  
  export const predefinedThemes: Record<ThemeName, Theme> = {
      light: { backgroundColor: "#FFFFFF", fontColor: "#000000", accentColor: "neutral-400" },
      dark: { backgroundColor: "#2D0A31", fontColor: "#FFFFFF", accentColor: "white" },
      black: { backgroundColor: "#000000", fontColor: "#FFFFFF", accentColor: "white" },
      white: { backgroundColor: "#FFFFFF", fontColor: "#000000", accentColor: "amber-800" },
      ocean: { backgroundColor: "#2193b0", fontColor: "#FFFFFF", accentColor: "white" },
      sunset: { backgroundColor: "#FF5841", fontColor: "#FFFFFF", accentColor: "white" },
      forest: { backgroundColor: "#228B22", fontColor: "#FFFFFF", accentColor: "white" },
      tech: { backgroundColor: "#1D1842", fontColor: "#FFFFFF", accentColor: "white" },
      cream: { backgroundColor: "#F2F0EA", fontColor: "#333333", accentColor: "white" },
      deepBlue: { backgroundColor: "#000428", fontColor: "#FFFFFF", accentColor: "white" },
      mint: { backgroundColor: "#00b09b", fontColor: "#FFFFFF", accentColor: "white" },
      purple: { backgroundColor: "#7C3AED", fontColor: "#FFFFFF", accentColor: "white" },
      pinkTeal: { backgroundColor: "#FF78AC", fontColor: "#FFFFFF", accentColor: "white" },
      sunsetOrange: { backgroundColor: "#FFAB00", fontColor: "#333333", accentColor: "white" },
      deepPurple: { backgroundColor: "#8E0D3C", fontColor: "#FFFFFF", accentColor: "white" },
      vibrantYellow: { backgroundColor: "#FFD43A", fontColor: "#582C12", accentColor: "white" },
      neonGreen: { backgroundColor: "#00D494", fontColor: "#FFFFFF", accentColor: "white" }
  };

export function getThemeNameFromTheme(theme: Theme | string): ThemeName {
    // If it's a Theme object, try to match it with predefined themes
    if (typeof theme === 'object' && theme !== null) {
        const matchingTheme = Object.entries(predefinedThemes).find(([_, value]) => 
            value.backgroundColor === theme.backgroundColor && 
            value.fontColor === theme.fontColor
        );
        return (matchingTheme?.[0] as ThemeName) || 'light'; // fallback to light
    }

    // If it's a string and valid ThemeName, use it
    if (typeof theme === 'string' && theme in predefinedThemes) {
        return theme as ThemeName;
    }

    // Default fallback
    return 'dark';
}

export function getThemeFromThemeName(themeName: ThemeName | string): Theme {
  // If it's a valid theme name, return the corresponding theme
  if (themeName in predefinedThemes) {
      return predefinedThemes[themeName as ThemeName];
  }

  // Default fallback
  return predefinedThemes.dark;
}

// Optional: Combined function to get Theme from either Theme or ThemeName
export function getTheme(theme: Theme | ThemeName | string): Theme {
  if (typeof theme === 'object' && theme !== null) {
      // If it's already a Theme object, validate and return it or fallback
      if ('backgroundColor' in theme && 'fontColor' in theme) {
          return theme;
      }
  }

  // Convert to ThemeName first, then get Theme
  const themeName = getThemeNameFromTheme(theme);
  return predefinedThemes[themeName];
}