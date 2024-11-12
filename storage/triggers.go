package storage

import (
	"gorm.io/gorm"
)

func (db *storage) beforeProductInsertTrigger() error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DROP TRIGGER IF EXISTS adjust_price_on_insert").Error; err != nil {
			return err
		}
		return tx.Exec(
			`
    CREATE TRIGGER adjust_price_on_insert
    BEFORE INSERT ON products
    FOR EACH ROW
    BEGIN
        -- Netherlands: Add 21% VAT
        IF NEW.Country = 'Netherlands' THEN
            SET NEW.Price = NEW.Price * 1.21;

        -- Germany: Add 19% VAT
        ELSEIF NEW.Country = 'Germany' THEN
            SET NEW.Price = NEW.Price * 1.19;

        -- France: Add 20% VAT
        ELSEIF NEW.Country = 'France' THEN
            SET NEW.Price = NEW.Price * 1.20;

        -- Default VAT (e.g. 10%) for other countries
        ELSE
            SET NEW.Price = NEW.Price * 1.10;
        END IF;
    END;
    `,
		).Error
	})
}

func (db *storage) beforeProductUpdateTrigger() error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DROP TRIGGER IF EXISTS adjust_price_on_update").Error; err != nil {
			return err
		}
		return tx.Exec(
			`
    CREATE TRIGGER adjust_price_on_update
    BEFORE UPDATE ON products
    FOR EACH ROW
    BEGIN
        -- Check if Country or Price has changed
        IF NEW.Country != OLD.Country OR NEW.Price != OLD.Price THEN
            -- Netherlands: Add 21% VAT
            IF NEW.Country = 'Netherlands' THEN
                SET NEW.Price = NEW.Price * 1.21;

            -- Germany: Add 19% VAT
            ELSEIF NEW.Country = 'Germany' THEN
                SET NEW.Price = NEW.Price * 1.19;

            -- France: Add 20% VAT
            ELSEIF NEW.Country = 'France' THEN
                SET NEW.Price = NEW.Price * 1.20;

            -- Default VAT (e.g., 10%) for other countries
            ELSE
                SET NEW.Price = NEW.Price * 1.10;
            END IF;
        END IF;
    END;
    `,
		).Error
	})
}
