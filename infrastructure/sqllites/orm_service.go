package sqllites

import (
	"database/sql"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/orms"
	"steve.care/network/domain/orms/skeletons"
)

type ormService struct {
	hashAdapter hash.Adapter
	skeleton    skeletons.Skeleton
	txPtr       *sql.Tx
}

func createOrmService(
	hashAdapter hash.Adapter,
	skeleton skeletons.Skeleton,
	txPtr *sql.Tx,
) orms.Service {
	out := ormService{
		hashAdapter: hashAdapter,
		skeleton:    skeleton,
		txPtr:       txPtr,
	}

	return &out
}

// Init initializes the service
func (app *ormService) Init(name string) error {
	return nil
}

// Insert inserts an instance
func (app *ormService) Insert(ins orms.Instance, path []string) error {
	return nil
}

/*
func (app *ormService) insertRoots(
	ins orms.Instance,
	roots roots.Roots,
) error {
	list := roots.List()
	for _, oneRoot := range list {
		err := app.insertRoot(ins, oneRoot)
		if err != nil {
			fmt.Printf("\n%s\n", err.Error())
			continue
		}

		return nil
	}

	return errors.New("the instance could not be inserted using the provided schema")
}

func (app *ormService) insertRoot(
	ins orms.Instance,
	root roots.Root,
) error {
	name := root.Name()
	chains := root.Chains()
	return app.insertChains(ins, chains, name, root)
}

func (app *ormService) insertGroup(
	ins orms.Instance,
	group schema_groups.Group,
	parentName string,
	root roots.Root,
) error {
	name := group.Name()
	chains := group.Chains()
	updatedParentName := name
	if parentName != "" {
		updatedParentName = fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, name)
	}

	return app.insertChains(ins, chains, updatedParentName, root)
}

func (app *ormService) insertChains(
	ins orms.Instance,
	chains schema_groups.MethodChains,
	parentName string,
	root roots.Root,
) error {
	list := chains.List()
	for _, oneChain := range list {
		isInserted, err := app.insertChain(ins, oneChain, parentName, root)
		if err != nil {
			return err
		}

		if !isInserted {
			continue
		}

		return nil
	}

	return nil
}

func (app *ormService) insertChain(
	ins orms.Instance,
	chain schema_groups.MethodChain,
	parentName string,
	root roots.Root,
) (bool, error) {

	errorString := ""
	conditionMethodName := chain.Condition()
	retValue, err := callMethodsOnInstance([]string{
		conditionMethodName,
	}, ins, &errorString)
	if err != nil {
		return false, err
	}

	typeName := reflect.TypeOf(&ins).Elem().Name()
	if errorString != "" {
		str := fmt.Sprintf("there was an error while calling a method chain's condition method (name: %s) on a reference instance (type: %s), the error was: %s", conditionMethodName, typeName, errorString)
		return false, errors.New(str)
	}

	if boolValue, ok := retValue.(bool); ok {
		if !boolValue {
			return false, nil
		}

		errorString := ""
		retrieverMethodNames := chain.Retriever()
		retValue, err := callMethodsOnInstance(retrieverMethodNames, ins, &errorString)
		if err != nil {
			return false, err
		}

		if errorString != "" {
			str := fmt.Sprintf("there was an error while calling a method chain's value method (chain: %s) on a reference instance (type: %s), the error was: %s", strings.Join(retrieverMethodNames, ","), typeName, errorString)
			return false, errors.New(str)
		}

		if castedValue, ok := retValue.(orms.Instance); ok {
			element := chain.Element()
			err = app.insertElement(castedValue, element, parentName, root)
			if err != nil {
				return false, err
			}

			return true, nil
		}

		str := fmt.Sprintf("the returned value was expected to contain an orms.Instance instance")
		return false, errors.New(str)

	}

	str := fmt.Sprintf("there was an error while calling a method chain's condition method (name: %s) on a reference instance (type: %s), the returned value was expected to be a bool, but it was NOT", conditionMethodName, typeName)
	return false, errors.New(str)
}

func (app *ormService) insertElement(
	ins orms.Instance,
	element schema_groups.Element,
	parentName string,
	root roots.Root,
) error {
	if element.IsResource() {
		resource := element.Resource()
		return app.insertResource(ins, resource, parentName, root)
	}

	group := element.Group()
	return app.insertGroup(ins, group, parentName, root)
}

func (app *ormService) insertResource(
	ins orms.Instance,
	resource schema_resources.Resource,
	parentName string,
	root roots.Root,
) error {
	key := resource.Key()
	fieldNames := []string{
		key.Name(),
	}

	errorString := ""
	keyMethods := key.Methods()
	typeName := reflect.TypeOf(&ins).Elem().Name()
	retPkValue, err := callMethodsOnInstance(keyMethods.Retriever(), ins, &errorString)
	if err != nil {
		return err
	}

	if errorString != "" {
		str := fmt.Sprintf("there was an error while calling a field key method (names: %s) on a reference instance (type: %s), the error was: %s", strings.Join(keyMethods.Retriever(), ","), typeName, errorString)
		return errors.New(str)
	}

	fieldValues := []interface{}{
		retPkValue,
	}

	fieldValuePlaceHolders := []string{
		"?",
	}

	fieldsList := resource.Fields().List()
	for _, oneField := range fieldsList {
		errorString := ""
		methods := oneField.Methods()
		retValue, err := callMethodsOnInstance(methods.Retriever(), ins, &errorString)
		if err != nil {
			return err
		}

		if errorString != "" {
			str := fmt.Sprintf("there was an error while calling a field method (names: %s) on a reference instance (type: %s), the error was: %s", strings.Join(methods.Retriever(), ","), typeName, errorString)
			return errors.New(str)
		}

		if retValue == nil && !oneField.CanBeNil() {
			str := fmt.Sprintf("the field (resource: %s, name: %s) is nil but is set as 'cannot be nil' in the schema", resource.Name(), oneField.Name())
			return errors.New(str)
		}

		// do not set the field in the query if the value is nil:
		if retValue == nil {
			continue
		}

		typ := oneField.Type()
		if typ.IsDependency() {
			// call the depenency retriever on the instance:
			dependency := typ.Dependency()
			retriever := dependency.Retriever()
			errorString := ""
			retValue, err := callMethodsOnInstance([]string{retriever}, ins, &errorString)
			if err != nil {
				return err
			}

			if errorString != "" {
				str := fmt.Sprintf("there was an error while calling a field dependency method (name: %s) on a reference instance (type: %s), the error was: %s", retriever, typeName, errorString)
				return errors.New(str)
			}

			// fetch the resource:
			groupNames := dependency.Groups()
			resourceName := dependency.Resource()
			path := append(groupNames, resourceName)
			retResourceSchema, err := root.Search(path)
			if err != nil {
				return err
			}

			// generate the parent name:
			dependencyParentName := strings.Join(groupNames, groupNameDelimiterForTableNames)

			// insert the resource value:
			if castedValue, ok := retValue.(orms.Instance); ok {
				err = app.insertResource(castedValue, retResourceSchema, dependencyParentName, root)
				if err != nil {
					return err
				}
			} else {
				str := fmt.Sprintf("the returned value was expected to contain an orms.Instance instance")
				return errors.New(str)
			}
		}

		fieldNames = append(fieldNames, oneField.Name())
		fieldValues = append(fieldValues, retValue)
		fieldValuePlaceHolders = append(fieldValuePlaceHolders, "?")
	}

	fieldNamesStr := strings.Join(fieldNames, ", ")
	fieldValuePlaceHoldersStr := strings.Join(fieldValuePlaceHolders, ", ")
	tableName := fmt.Sprintf("%s%s%s", parentName, groupNameDelimiterForTableNames, resource.Name())
	queryStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, fieldNamesStr, fieldValuePlaceHoldersStr)
	_, err = app.txPtr.Exec(queryStr, fieldValues...)
	if err != nil {
		return err
	}

	return nil
}*/

// Delete deletes an instance
func (app *ormService) Delete(path []string, hash hash.Hash) error {
	return nil
}
