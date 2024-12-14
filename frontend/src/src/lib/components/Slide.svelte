<script>
    import { onMount } from 'svelte';
    
    const metrics = {
      mau: {
        start: '300K',
        jan2024: '1M',
        current: '3.3M',
        goal: 'XXX'
      },
      roas: {
        current: '190%*',
        target: '200%+'
      }
    };
  </script>
  
  <div class="slide">
    <div class="header">
      <h1>Яндекс Пэй: Стратегия развития платного продвижения 2025</h1>
      <div class="subtitle">Performance Marketing Strategy</div>
    </div>
  
    <div class="content">
      <div class="achievements">
        <div class="metric-card">
          <span class="label">MAU Growth</span>
          <div class="value-row">
            <span class="value-item">
              <small>Sep 2023</small>
              {metrics.mau.start}
            </span>
            <span class="arrow">→</span>
            <span class="value-item">
              <small>Jan 2024</small>
              {metrics.mau.jan2024}
            </span>
            <span class="arrow">→</span>
            <span class="value-item highlight">
              <small>Current</small>
              {metrics.mau.current}
            </span>
            <span class="arrow">→</span>
            <span class="value-item target">
              <small>2025 Goal</small>
              {metrics.mau.goal}
            </span>
          </div>
        </div>
        
        <div class="metric-card">
          <span class="label">Performance Metrics</span>
          <div class="preliminary-notice">
            * Предварительные данные по всей базе пользователей
          </div>
          <div class="metrics-grid">
            <div class="metric">
              <span>ROAS</span>
              <strong class="highlight">190%*</strong>
            </div>
            <div class="metric">
              <span>LTV/CAC</span>
              <strong class="highlight">1.9x*</strong>
            </div>
          </div>
        </div>
      </div>
  
      <div class="strategy-section">
        <h2>Стратегия 2025</h2>
        
        <div class="strategy-grid">
          <div class="strategy-card">
            <div class="phase">Phase 1: Q1 2025</div>
            <div class="live-experiment">
              <span class="pulse"></span>
              Эксперимент активен
            </div>
            <h3>Валидация метрик</h3>
            <ul>
              <li>Подтверждение LTV платных пользователей</li>
              <li>Измерение уровня каннибализации</li>
              <li>Расчет unit-экономики по каналам</li>
            </ul>
            <div class="timeline-note">
              Анализ результатов: начало января 2025
            </div>
          </div>
  
          <div class="strategy-card">
            <div class="phase">Phase 2: Q2-Q4 2025</div>
            <h3>Масштабирование</h3>
            <ul>
              <li>Positive ROI каналы</li>
              <li>CRM оптимизация</li>
              <li>Таргетинг premium сегментов</li>
            </ul>
          </div>
        </div>
      </div>
  
      <div class="risks-section">
        <h2>Ключевые факторы успеха</h2>
        <div class="risks-grid">
          <div class="risk-item">
            <div class="risk-header">
              <span class="icon">📊</span>
              <span>Точность атрибуции</span>
            </div>
            <p class="risk-description">
              Корректное определение источника привлечения пользователя для точного расчета эффективности каналов
            </p>
          </div>
          <div class="risk-item">
            <div class="risk-header">
              <span class="icon">📈</span>
              <span>Контроль каннибализации</span>
            </div>
            <p class="risk-description">
              Минимизация замещения органического трафика платным для оптимизации маркетингового бюджета
            </p>
          </div>
          <div class="risk-item">
            <div class="risk-header">
              <span class="icon">💰</span>
              <span>Устойчивый ROAS</span>
            </div>
            <p class="risk-description">
              Поддержание положительного возврата инвестиций при масштабировании рекламных кампаний
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
  
  <style>
    .slide {
      background: #0A0A0A;
      color: #FFFFFF;
      padding: 2rem;
      font-family: 'SF Pro Display', -apple-system, sans-serif;
      /* height: 100vh; */
      display: flex;
      flex-direction: column;
    }
  
    .header {
      text-align: center;
      margin-bottom: 2rem;
    }
  
    h1 {
      font-size: 2.5rem;
      margin: 0;
      background: linear-gradient(90deg, #FF5C5C, #FF8C8C);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
    }
  
    .subtitle {
      color: #666;
      font-size: 1.2rem;
      margin-top: 0.5rem;
    }
  
    .content {
      display: flex;
      flex-direction: column;
      gap: 2rem;
    }
  
    .achievements {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 1rem;
    }
  
    .metric-card {
      background: rgba(255,255,255,0.05);
      border-radius: 12px;
      padding: 1.5rem;
      border: 1px solid rgba(255,255,255,0.1);
    }
  
    .label {
      color: #666;
      font-size: 0.9rem;
      text-transform: uppercase;
      letter-spacing: 1px;
    }
  
    .value-row {
      display: flex;
      align-items: center;
      gap: 1rem;
      margin-top: 1rem;
      flex-wrap: wrap;
      justify-content: space-between;
    }
  
    .value-item {
      display: flex;
      flex-direction: column;
      align-items: center;
    }
  
    .value-item small {
      color: #666;
      font-size: 0.8rem;
    }
  
    .highlight {
      color: #FF5C5C;
    }
  
    .arrow {
      color: #666;
    }
  
    .strategy-grid {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 1rem;
      margin-top: 1rem;
    }
  
    .strategy-card {
      background: rgba(255,255,255,0.05);
      border-radius: 12px;
      padding: 1.5rem;
      border: 1px solid rgba(255,255,255,0.1);
    }
  
    .phase {
      color: #FF5C5C;
      font-size: 0.9rem;
      margin-bottom: 0.5rem;
    }
  
    h2 {
      font-size: 1.5rem;
      margin: 0;
      margin-bottom: 1rem;
    }
  
    h3 {
      font-size: 1.2rem;
      margin: 0;
      margin-bottom: 1rem;
    }
  
    ul {
      margin: 0;
      padding-left: 1.2rem;
      color: #CCC;
    }
  
    .risks-grid {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      gap: 1rem;
    }
  
    .risk-item {
      background: rgba(255,255,255,0.05);
      border-radius: 8px;
      padding: 1.5rem;
      display: flex;
      flex-direction: column;
      align-items: flex-start;
    }
  
    .risk-header {
      display: flex;
      align-items: center;
      gap: 0.5rem;
      margin-bottom: 0.8rem;
    }
  
    .icon {
      font-size: 1.5rem;
    }
  
    .metrics-grid {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 1rem;
      margin-top: 1rem;
    }
  
    .metric {
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 0.5rem;
    }
  
    .metric span {
      color: #666;
    }
  
    .metric strong {
      font-size: 1.5rem;
    }
  
    .preliminary-notice {
      color: #666;
      font-size: 0.8rem;
      font-style: italic;
      margin: 0.5rem 0;
    }
  
    .live-experiment {
      display: flex;
      align-items: center;
      gap: 0.5rem;
      color: #4CAF50;
      margin: 0.5rem 0;
    }
  
    .pulse {
      width: 10px;
      height: 10px;
      background-color: #4CAF50;
      border-radius: 50%;
      animation: pulse 2s infinite;
    }
  
    @keyframes pulse {
      0% {
        transform: scale(0.95);
        box-shadow: 0 0 0 0 rgba(76, 175, 80, 0.7);
      }
      70% {
        transform: scale(1);
        box-shadow: 0 0 0 10px rgba(76, 175, 80, 0);
      }
      100% {
        transform: scale(0.95);
        box-shadow: 0 0 0 0 rgba(76, 175, 80, 0);
      }
    }
  
    .timeline-note {
      margin-top: 1rem;
      font-size: 0.9rem;
      color: #FF5C5C;
      font-style: italic;
    }
  
    .risk-description {
      margin: 0;
      font-size: 0.9rem;
      color: #999;
      line-height: 1.4;
    }
  </style>
  