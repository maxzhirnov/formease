export type AnimationPattern = 'pattern1' | 'pattern2' | 'pattern3';

export interface Shape {
  width: number;
  height: number;
  color: string;
  position: { top?: string; left?: string; right?: string; bottom?: string };
  blur: number;
  opacity: number;
  animationDuration: number;
  animationPattern: AnimationPattern;
}

export type FloatingShapesTheme = 'default' | 'purple' | 'ocean' | 'sunset' | 'forest' | 'neon' | 'pastel' | 'monochrome' | 'autumn' | 'spring';

const themes: Record<FloatingShapesTheme, Shape[]> = {
    default: [
      {
        width: 600,
        height: 600,
        color: "#FF3E8D",
        position: { top: "-10%", left: "-5%" },
        blur: 80,
        opacity: 0.5,
        animationDuration: 20,
        animationPattern: "pattern1"
      }
    ],
    purple: [
      {
        width: 600,
        height: 600,
        color: "#7C3AED",
        position: { top: "-10%", right: "5%" },
        blur: 100,
        opacity: 0.4,
        animationDuration: 20,
        animationPattern: "pattern1"
      },
      {
        width: 500,
        height: 500,
        color: "#9333EA",
        position: { bottom: "-5%", left: "10%" },
        blur: 90,
        opacity: 0.6,
        animationDuration: 25,
        animationPattern: "pattern2"
      }
    ],
    ocean: [
      {
        width: 600,
        height: 600,
        color: "#03fc98",
        position: { top: "-10%", right: "5%" },
        blur: 100,
        opacity: 0.4,
        animationDuration: 20,
        animationPattern: "pattern1"
      },
      {
        width: 500,
        height: 500,
        color: "#3B82F6",
        position: { bottom: "-5%", left: "10%" },
        blur: 90,
        opacity: 0.6,
        animationDuration: 25,
        animationPattern: "pattern2"
      }
    ],
    sunset: [
      {
        width: 600,
        height: 600,
        color: "#F59E0B",
        position: { top: "-10%", right: "5%" },
        blur: 100,
        opacity: 0.4,
        animationDuration: 20,
        animationPattern: "pattern1"
      },
      {
        width: 500,
        height: 500,
        color: "#EF4444",
        position: { bottom: "-5%", left: "10%" },
        blur: 90,
        opacity: 0.6,
        animationDuration: 25,
        animationPattern: "pattern2"
      }
    ],
    forest: [
      {
        width: 700,
        height: 700,
        color: "#059669",
        position: { top: "-15%", left: "-10%" },
        blur: 120,
        opacity: 0.3,
        animationDuration: 30,
        animationPattern: "pattern3"
      },
      {
        width: 550,
        height: 550,
        color: "#10B981",
        position: { bottom: "-10%", right: "5%" },
        blur: 100,
        opacity: 0.5,
        animationDuration: 25,
        animationPattern: "pattern1"
      }
    ],
    neon: [
      {
        width: 400,
        height: 400,
        color: "#FF00FF",
        position: { top: "10%", left: "20%" },
        blur: 50,
        opacity: 0.7,
        animationDuration: 15,
        animationPattern: "pattern2"
      },
      {
        width: 300,
        height: 300,
        color: "#00FFFF",
        position: { bottom: "15%", right: "25%" },
        blur: 40,
        opacity: 0.6,
        animationDuration: 18,
        animationPattern: "pattern3"
      }
    ],
    pastel: [
      {
        width: 650,
        height: 650,
        color: "#FFC0CB",
        position: { top: "-5%", left: "-5%" },
        blur: 110,
        opacity: 0.3,
        animationDuration: 28,
        animationPattern: "pattern1"
      },
      {
        width: 450,
        height: 450,
        color: "#87CEFA",
        position: { bottom: "-8%", right: "15%" },
        blur: 80,
        opacity: 0.4,
        animationDuration: 22,
        animationPattern: "pattern2"
      }
    ],
    monochrome: [
      {
        width: 550,
        height: 550,
        color: "#808080",
        position: { top: "-8%", right: "-5%" },
        blur: 90,
        opacity: 0.3,
        animationDuration: 26,
        animationPattern: "pattern3"
      },
      {
        width: 400,
        height: 400,
        color: "#A9A9A9",
        position: { bottom: "-10%", left: "8%" },
        blur: 70,
        opacity: 0.5,
        animationDuration: 20,
        animationPattern: "pattern1"
      }
    ],
    autumn: [
      {
        width: 600,
        height: 600,
        color: "#FFA500",
        position: { top: "-12%", left: "-8%" },
        blur: 95,
        opacity: 0.4,
        animationDuration: 24,
        animationPattern: "pattern2"
      },
      {
        width: 500,
        height: 500,
        color: "#8B4513",
        position: { bottom: "-7%", right: "10%" },
        blur: 85,
        opacity: 0.5,
        animationDuration: 28,
        animationPattern: "pattern3"
      }
    ],
    spring: [
      {
        width: 550,
        height: 550,
        color: "#FF69B4",
        position: { top: "-5%", right: "-3%" },
        blur: 75,
        opacity: 0.4,
        animationDuration: 22,
        animationPattern: "pattern1"
      },
      {
        width: 450,
        height: 450,
        color: "#98FB98",
        position: { bottom: "-6%", left: "12%" },
        blur: 65,
        opacity: 0.5,
        animationDuration: 26,
        animationPattern: "pattern2"
      }
    ]
  };

export { themes };
