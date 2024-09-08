package system

func Migrations() error {
	var err error
	_, err = Db.Exec(`
				create table if not exists favorites (
						mal_id integer primary key not null,
						created datetime not null default current_timestamp,
						updated datetime not null default current_timestamp,
						title text not null,
						image text not null
				)`)
	if err != nil {
		return err
	}
	return nil
}
