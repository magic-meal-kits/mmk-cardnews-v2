import React from "react";
import { AbsoluteFill, Img, interpolate, staticFile, useCurrentFrame } from "remotion";
import { POPEYE } from "../magazine-styles";

interface CardProps {
  durationInFrames: number;
}

export const Card02: React.FC<CardProps> = ({ durationInFrames }) => {
  const frame = useCurrentFrame();

  const FADE_IN = 9;
  const TITLE_IN = 15;
  const BODY_IN = 24;
  const FADE_OUT = 15;

  const imgOpacity = interpolate(frame, [0, FADE_IN], [0, 1], { extrapolateRight: "clamp" });
  const tagOpacity = interpolate(frame, [TITLE_IN - 6, TITLE_IN + 6], [0, 1], { extrapolateRight: "clamp" });
  const titleOpacity = interpolate(frame, [TITLE_IN, TITLE_IN + 12], [0, 1], { extrapolateRight: "clamp" });
  const bodyOpacity = interpolate(frame, [BODY_IN, BODY_IN + 12], [0, 1], { extrapolateRight: "clamp" });
  const fadeOut = interpolate(frame, [durationInFrames - FADE_OUT, durationInFrames], [1, 0], { extrapolateLeft: "clamp" });

  return (
    <AbsoluteFill style={{ display: "flex", flexDirection: "column", opacity: fadeOut }}>
      {/* Image top 55% */}
      <div style={{ width: "100%", height: "55%", overflow: "hidden", position: "relative" }}>
        <Img
          src={staticFile("card-02.png")}
          style={{ width: "100%", height: "100%", objectFit: "cover", opacity: imgOpacity }}
        />
        <div
          style={{
            position: "absolute",
            inset: 0,
            background: "linear-gradient(to bottom, rgba(0,0,0,0.05) 0%, rgba(0,0,0,0.25) 100%)",
          }}
        />
      </div>

      {/* Text bottom 45% */}
      <div
        style={{
          width: "100%",
          height: "45%",
          padding: "48px 72px 90px 72px",
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          background: POPEYE.colors.background,
          position: "relative",
        }}
      >
        <div
          style={{
            fontFamily: "'DM Sans', sans-serif",
            fontWeight: 700,
            fontSize: 30,
            letterSpacing: "0.12em",
            textTransform: "uppercase",
            color: "#D4563A",
            marginBottom: 24,
            opacity: tagOpacity,
          }}
        >
          카구라자카 노포 3곳
        </div>

        {/* 시마킨 */}
        <div style={{ opacity: titleOpacity, marginBottom: 20 }}>
          <h2
            style={{
              fontFamily: "'Noto Serif KR', serif",
              fontWeight: 700,
              fontSize: 48,
              lineHeight: 1.2,
              color: POPEYE.colors.primary,
              marginBottom: 6,
            }}
          >
            시마킨 (志満金)
          </h2>
          <p
            style={{
              fontFamily: "'Noto Sans KR', sans-serif",
              fontWeight: 400,
              fontSize: 32,
              color: "#B5A89A",
              letterSpacing: "0.02em",
            }}
          >
            1869년 창업 · 장어 요리 · 6대째
          </p>
        </div>

        {/* 키노젠 + 우오토쿠 */}
        <div style={{ opacity: bodyOpacity, display: "flex", flexDirection: "column", gap: 14 }}>
          <div>
            <span
              style={{
                fontFamily: "'Noto Serif KR', serif",
                fontWeight: 700,
                fontSize: 40,
                color: POPEYE.colors.primary,
              }}
            >
              키노젠
            </span>
            <span
              style={{
                fontFamily: "'Noto Sans KR', sans-serif",
                fontWeight: 400,
                fontSize: 32,
                color: POPEYE.colors.muted,
                marginLeft: 16,
              }}
            >
              1860년 창업 · 말차 바바루아 원조
            </span>
          </div>
          <div>
            <span
              style={{
                fontFamily: "'Noto Serif KR', serif",
                fontWeight: 700,
                fontSize: 40,
                color: POPEYE.colors.primary,
              }}
            >
              우오토쿠
            </span>
            <span
              style={{
                fontFamily: "'Noto Sans KR', sans-serif",
                fontWeight: 400,
                fontSize: 32,
                color: POPEYE.colors.muted,
                marginLeft: 16,
              }}
            >
              1920년 창업 · 가이세키 · 6대째
            </span>
          </div>
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
            color: "rgba(0,0,0,0.3)",
          }}
        >
          2 / 3
        </div>
      </div>
    </AbsoluteFill>
  );
};
