package test

import s "github.com/twitchscience/scoop_protocol/schema"

func EventTest1() s.Event {
	return s.Event{
		EventName: "test_event_1",
		Version:   1,
		TableOption: s.TableOption{
			DistKey: []string{"test_event_1_outbound_col_1"},
			SortKey: []string{},
		},
		Columns: []s.ColumnDefinition{
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "test_event_1_inbound_col_1",
				OutboundName:          "test_event_1_outbound_col_1",
				ColumnCreationOptions: "(500)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "test_event_1_inbound_col_2",
				OutboundName:          "test_event_1_outbound_col_2",
				ColumnCreationOptions: "(500)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "test_event_1_inbound_col_3",
				OutboundName:          "test_event_1_outbound_col_3",
				ColumnCreationOptions: "(500)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "test_event_1_inbound_col_4",
				OutboundName:          "test_event_1_outbound_col_4",
				ColumnCreationOptions: "(500)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "test_event_1_inbound_col_5",
				OutboundName:          "test_event_1_outbound_col_5",
				ColumnCreationOptions: "(500)",
			},
		},
		ParentMigration: s.Migration{},
	}
}

func EventTest1Empty() s.Event {
	return s.Event{
		EventName:       "test_event_1",
		Version:         1,
		ParentMigration: s.Migration{},
	}
}

func Migration1OnEvent1() s.Migration {
	return s.Migration{
		TableOperation: "add",
		Name:           "test_event_1",
		TableOption: s.TableOption{
			DistKey: []string{"test_event_1_new_outbound_col_1"},
			SortKey: []string{},
		},
		ColumnOperations: []s.ColumnOperation{
			s.ColumnOperation{
				Operation:    "add",
				InboundName:  "test_event_1_new_inbound_col_1",
				OutboundName: "test_event_1_new_outbound_col_1",
				NewColumnDefinition: s.ColumnDefinition{
					Transformer:           "varchar",
					InboundName:           "test_event_1_new_inbound_col_1",
					OutboundName:          "test_event_1_new_outbound_col_1",
					ColumnCreationOptions: "(500)",
				},
			},
			s.ColumnOperation{
				Operation:    "add",
				InboundName:  "test_event_1_new_inbound_col_2",
				OutboundName: "test_event_1_new_outbound_col_2",
				NewColumnDefinition: s.ColumnDefinition{
					Transformer:           "varchar",
					InboundName:           "test_event_1_new_inbound_col_2",
					OutboundName:          "test_event_1_new_outbound_col_2",
					ColumnCreationOptions: "(500)",
				},
			},
			s.ColumnOperation{
				Operation:    "add",
				InboundName:  "test_event_1_new_inbound_col_3",
				OutboundName: "test_event_1_new_outbound_col_3",
				NewColumnDefinition: s.ColumnDefinition{
					Transformer:           "varchar",
					InboundName:           "test_event_1_new_inbound_col_3",
					OutboundName:          "test_event_1_new_outbound_col_3",
					ColumnCreationOptions: "(500)",
				},
			},
			s.ColumnOperation{
				Operation:    "add",
				InboundName:  "test_event_1_new_inbound_col_4",
				OutboundName: "test_event_1_new_outbound_col_4",
				NewColumnDefinition: s.ColumnDefinition{
					Transformer:           "varchar",
					InboundName:           "test_event_4_new_inbound_col_4",
					OutboundName:          "test_event_4_new_outbound_col_4",
					ColumnCreationOptions: "(500)",
				},
			},
		},
	}
}

func Migration2OnEvent1() s.Migration {
	return s.Migration{
		TableOperation:   "remove",
		Name:             "test_event_1",
		TableOption:      s.TableOption{},
		ColumnOperations: []s.ColumnOperation{},
	}
}

func Migration3OnEvent1() s.Migration {
	{
		return s.Migration{
			TableOperation: "update",
			Name:           "test_event_1",
			TableOption: s.TableOption{
				DistKey: []string{"test_event_1_outbound_col_1"},
				SortKey: []string{},
			},
			ColumnOperations: []s.ColumnOperation{
				s.ColumnOperation{
					Operation:    "add",
					InboundName:  "test_event_1_new_inbound_col_1",
					OutboundName: "test_event_1_new_outbound_col_1",
					NewColumnDefinition: s.ColumnDefinition{
						Transformer:           "varchar",
						InboundName:           "test_event_1_new_inbound_col_1",
						OutboundName:          "test_event_1_new_outbound_col_1",
						ColumnCreationOptions: "(500)",
					},
				},
				s.ColumnOperation{
					Operation:    "update",
					InboundName:  "test_event_1_inbound_col_2",
					OutboundName: "test_event_1_outbound_col_2",
					NewColumnDefinition: s.ColumnDefinition{
						Transformer:           "varchar",
						InboundName:           "test_event_1_new_inbound_col_2_updated",
						OutboundName:          "test_event_1_new_outbound_col_2_updated",
						ColumnCreationOptions: "(500)",
					},
				},
				s.ColumnOperation{
					Operation:           "remove",
					InboundName:         "test_event_1_inbound_col_3",
					OutboundName:        "test_event_1_outbound_col_3",
					NewColumnDefinition: s.ColumnDefinition{},
				},
			},
		}
	}
}

func SimEvent1Version1() s.Event {
	return s.Event{}
}
func SimEvent1Migration1() s.Migration {
	return s.Migration{}
}

func SimEvent1Version2() s.Event {
	return s.Event{}
}
func SimEvent1Migration2() s.Migration {
	return s.Migration{}
}

func SimEvent1Version3() s.Event {
	return s.Event{}
}
func SimEvent1Migration3() s.Migration {
	return s.Migration{}
}

func SimEvent1Version4() s.Event {
	return s.Event{}
}
func SimEvent1Migration4() s.Migration {
	return s.Migration{}
}

func SimEvent1Version5() s.Event {
	return s.Event{
		EventName: "chat_ignore_client",
		Version:   5,
		TableOption: s.TableOption{
			DistKey: []string{"device_id"},
			SortKey: []string{},
		},
		Columns: []s.ColumnDefinition{
			s.ColumnDefinition{
				Transformer:           "f@timestamp@unix",
				InboundName:           "time",
				OutboundName:          "time",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "ip",
				OutboundName:          "ip",
				ColumnCreationOptions: "(15)",
			},
			s.ColumnDefinition{
				Transformer:           "ipCity",
				InboundName:           "ip",
				OutboundName:          "city",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipCountry",
				InboundName:           "ip",
				OutboundName:          "country",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipRegion",
				InboundName:           "ip",
				OutboundName:          "region",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipAsnInteger",
				InboundName:           "ip",
				OutboundName:          "asn_id",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "device_id",
				OutboundName:          "device_id",
				ColumnCreationOptions: "(32)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "ui_context",
				OutboundName:          "ui_context",
				ColumnCreationOptions: "(25)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "from_login",
				OutboundName:          "from_login",
				ColumnCreationOptions: "(25)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "ignored_login",
				OutboundName:          "ignored_login",
				ColumnCreationOptions: "(25)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "reason",
				OutboundName:          "reason",
				ColumnCreationOptions: "(32)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "platform",
				OutboundName:          "platform",
				ColumnCreationOptions: "(32)",
			},
		},
		ParentMigration: s.Migration{},
	}
}

func SimEvent2Version1() s.Event {
	return s.Event{}
}
func SimEvent2Migration1() s.Migration {
	return s.Migration{}
}

func SimEvent2Version2() s.Event {
	return s.Event{
		EventName: "login_success",
		Version:   3,
		TableOption: s.TableOption{
			DistKey: []string{"device_id"},
			SortKey: []string{},
		},
		Columns: []s.ColumnDefinition{
			s.ColumnDefinition{
				Transformer:           "f@timestamp@unix",
				InboundName:           "time",
				OutboundName:          "time",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "ip",
				OutboundName:          "ip",
				ColumnCreationOptions: "(15)",
			},
			s.ColumnDefinition{
				Transformer:           "ipCity",
				InboundName:           "ip",
				OutboundName:          "city",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipCountry",
				InboundName:           "ip",
				OutboundName:          "country",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipRegion",
				InboundName:           "ip",
				OutboundName:          "region",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipAsnInteger",
				InboundName:           "ip",
				OutboundName:          "asn_id",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "device_id",
				OutboundName:          "device_id",
				ColumnCreationOptions: "(32)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "login",
				OutboundName:          "login",
				ColumnCreationOptions: "(16)",
			},
		},
		ParentMigration: s.Migration{},
	}
}
func SimEvent2Migration2() s.Migration {
	return s.Migration{}
}

func SimEvent2Version3() s.Event {
	return s.Event{
		EventName: "login_success",
		Version:   3,
		TableOption: s.TableOption{
			DistKey: []string{"device_id"},
			SortKey: []string{},
		},
		Columns: []s.ColumnDefinition{
			s.ColumnDefinition{
				Transformer:           "f@timestamp@unix",
				InboundName:           "time",
				OutboundName:          "time",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "ip",
				OutboundName:          "ip",
				ColumnCreationOptions: "(15)",
			},
			s.ColumnDefinition{
				Transformer:           "ipCity",
				InboundName:           "ip",
				OutboundName:          "city",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipCountry",
				InboundName:           "ip",
				OutboundName:          "country",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipRegion",
				InboundName:           "ip",
				OutboundName:          "region",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipAsnInteger",
				InboundName:           "ip",
				OutboundName:          "asn_id",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "device_id",
				OutboundName:          "device_id",
				ColumnCreationOptions: "(32)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "login",
				OutboundName:          "login",
				ColumnCreationOptions: "(25)",
			},
		},
		ParentMigration: s.Migration{},
	}
}
func SimEvent2Migration3() s.Migration {
	return s.Migration{}
}

func SimEvent2Version4() s.Event {
	return s.Event{
		EventName: "login_success",
		Version:   4,
		TableOption: s.TableOption{
			DistKey: []string{"device_id"},
			SortKey: []string{},
		},
		Columns: []s.ColumnDefinition{
			s.ColumnDefinition{
				Transformer:           "f@timestamp@unix",
				InboundName:           "time",
				OutboundName:          "time",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "ip",
				OutboundName:          "ip",
				ColumnCreationOptions: "(15)",
			},
			s.ColumnDefinition{
				Transformer:           "ipCity",
				InboundName:           "ip",
				OutboundName:          "city",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipCountry",
				InboundName:           "ip",
				OutboundName:          "country",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipRegion",
				InboundName:           "ip",
				OutboundName:          "region",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "ipAsnInteger",
				InboundName:           "ip",
				OutboundName:          "asn_id",
				ColumnCreationOptions: "",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "device_id",
				OutboundName:          "device_id",
				ColumnCreationOptions: "(32)",
			},
			s.ColumnDefinition{
				Transformer:           "varchar",
				InboundName:           "login",
				OutboundName:          "login",
				ColumnCreationOptions: "(25)",
			},
			s.ColumnDefinition{
				Transformer:           "bool",
				InboundName:           "two_factor_enabled",
				OutboundName:          "two_factor_enabled",
				ColumnCreationOptions: "",
			},
		},
		ParentMigration: s.Migration{},
	}
}
