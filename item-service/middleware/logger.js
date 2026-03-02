const config = require('../config');

/**
 * Structured logging utility
 */
const logger = {
  _format(level, message, meta = {}) {
    const logEntry = {
      timestamp: new Date().toISOString(),
      level,
      service: config.service.name,
      message,
      ...meta,
    };

    if (config.logging.format === 'json') {
      return JSON.stringify(logEntry);
    }

    // Simple format for development
    const metaStr = Object.keys(meta).length > 0 ? ` ${JSON.stringify(meta)}` : '';
    return `[${logEntry.timestamp}] ${level.toUpperCase()} [${config.service.name}]: ${message}${metaStr}`;
  },

  info(message, meta) {
    console.log(this._format('info', message, meta));
  },

  error(message, meta) {
    console.error(this._format('error', message, meta));
  },

  warn(message, meta) {
    console.warn(this._format('warn', message, meta));
  },

  debug(message, meta) {
    if (config.logging.level === 'debug') {
      console.log(this._format('debug', message, meta));
    }
  },
};

/**
 * HTTP Request logging middleware
 */
const requestLogger = (req, res, next) => {
  const start = Date.now();

  // Log when response finishes
  res.on('finish', () => {
    const duration = Date.now() - start;
    const logData = {
      method: req.method,
      path: req.originalUrl,
      statusCode: res.statusCode,
      duration: `${duration}ms`,
      ip: req.ip || req.connection.remoteAddress,
    };

    // Choose log level based on status code
    if (res.statusCode >= 500) {
      logger.error('Request failed', logData);
    } else if (res.statusCode >= 400) {
      logger.warn('Request error', logData);
    } else {
      logger.info('Request completed', logData);
    }
  });

  next();
};

module.exports = { logger, requestLogger };
