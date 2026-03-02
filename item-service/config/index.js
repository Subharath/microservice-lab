require('dotenv').config();

const config = {
  // Service info
  service: {
    name: 'item-service',
    version: process.env.SERVICE_VERSION || '1.0.0',
  },

  // Server config
  server: {
    port: parseInt(process.env.PORT, 10) || 3001,
    env: process.env.NODE_ENV || 'development',
  },

  // CORS config
  cors: {
    origin: process.env.CORS_ORIGIN || '*',
    credentials: true,
  },

  // Logging config
  logging: {
    level: process.env.LOG_LEVEL || 'info',
    format: process.env.LOG_FORMAT || 'json', // 'json' or 'simple'
  },

  // Other services URLs (for inter-service communication)
  services: {
    orderService: process.env.ORDER_SERVICE_URL || 'http://order-service:3002',
    paymentService: process.env.PAYMENT_SERVICE_URL || 'http://payment-service:3003',
  },
};

// Freeze config to prevent modifications
Object.freeze(config);

module.exports = config;
