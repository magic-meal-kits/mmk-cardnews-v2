import React from "react";
import { AbsoluteFill, Img, interpolate, staticFile, useCurrentFrame } from "remotion";
import { POPEYE } from "../magazine-styles";

interface CardProps {
  durationInFrames: number;
}

export const Card01: React.FC<CardProps> = ({ durationInFrames }) => {
  const frame = useCurrentFrame();

  const FADE_IN = 9;
  const TITLE_IN = 15;
  const BODY_IN = 24;
  const FADE_OUT = 15;

  const bgOpacity = interpolate(frame, [0, FADE_IN], [0, 1], { extrapolateRight: "clamp" });
  const tagOpacity = interpolate(frame, [TITLE_IN - 6, TITLE_IN + 6], [0, 1], { extrapolateRight: "clamp" });
  const titleOpacity = interpolate(frame, [TITLE_IN, TITLE_IN + 12], [0, 1], { extrapolateRight: "clamp" });
  const titleY = interpolate(frame, [TITLE_IN, TITLE_IN + 12], [24, 0], { extrapolateRight: "clamp" });
  const bodyOpacity = interpolate(frame, [BODY_IN, BODY_IN + 12], [0, 1], { extrapolateRight: "clamp" });
  const fadeOut = interpolate(frame, [durationInFrames - FADE_OUT, durationInFrames], [1, 0], { extrapolateLeft: "clamp" });

  return (
    <AbsoluteFill style={{ opacity: fadeOut }}>
      <Img
        src={staticFile("card-01.png")}
        style={{ width: "100%", height: "100%", objectFit: "cover", opacity: bgOpacity }}
      />
      <div
        style={{
          position: "absolute",
          inset: 0,
          background: POPEYE.overlay.cover,
        }}
      />

      {/* Top label */}
      <div
        style={{
          position: "absolute",
          top: 72,
          left: 72,
          opacity: tagOpacity,
        }}
      >
        <div
          style={{
            fontFamily: "'DM Sans', sans-serif",
            fontWeight: 700,
            fontSize: 30,
            letterSpacing: "0.15em",
            textTransform: "uppercase",
            color: "rgba(255,255,255,0.65)",
          }}
        >
          TOKYO · KAGURAZAKA
        </div>
      </div>

      {/* Bottom text zone */}
      <div
        style={{
          position: "absolute",
          bottom: 100,
          left: 72,
          right: 72,
        }}
      >
        <div
          style={{
            fontFamily: "'DM Sans', sans-serif",
            fontWeight: 700,
            fontSize: 32,
            letterSpacing: "0.12em",
            textTransform: "uppercase",
            color: "#D4563A",
            marginBottom: 20,
            opacity: bodyOpacity,
          }}
        >
          노포 맛집 3곳
        </div>
        <h2
          style={{
            fontFamily: "'Noto Serif KR', serif",
            fontWeight: 700,
            fontSize: 56,
            lineHeight: 1.2,
            color: "#FAFAFA",
            opacity: titleOpacity,
            transform: `translateY(${titleY}px)`,
            marginBottom: 20,
            letterSpacing: "-0.01em",
          }}
        >
          도쿄의 골목,
          <br />
          100년을 버텨온 맛
        </h2>
        <p
          style={{
            fontFamily: "'Noto Sans KR', sans-serif",
            fontWeight: 400,
            fontSize: 36,
            lineHeight: 1.55,
            color: "rgba(250,250,250,0.85)",
            opacity: bodyOpacity,
          }}
        >
          카구라자카에서 시간이 멈춘 세 곳의 노포를 소개합니다
        </p>
      </div>

      {/* Page number */}
      <div
        style={{
          position: "absolute",
          bottom: 40,
          left: "50%",
          transform: "translateX(-50%)",
          fontSize: 28,
          fontWeight: 300,
          letterSpacing: 3,
          color: "rgba(255,255,255,0.4)",
        }}
      >
        1 / 3
      </div>
    </AbsoluteFill>
  );
};
