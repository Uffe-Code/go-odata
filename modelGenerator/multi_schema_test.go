package modelGenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var multiSchemaEdmxSchema = `<edmx:Edmx xmlns:edmx="http://docs.oasis-open.org/odata/ns/edmx" Version="4.0">
<edmx:DataServices>
<Schema xmlns="http://docs.oasis-open.org/odata/ns/edm" Namespace="Trippin.Model">
<EntityType Name="Person">
<Key>
<PropertyRef Name="UserName"/>
</Key>
<Property Name="UserName" Type="Edm.String" Nullable="false"/>
<Property Name="FirstName" Type="Edm.String" Nullable="false"/>
<Property Name="LastName" Type="Edm.String" MaxLength="26"/>
<Property Name="MiddleName" Type="Edm.String"/>
<Property Name="Gender" Type="Trippin.Model.PersonGender" Nullable="false"/>
<Property Name="Age" Type="Edm.Int64"/>
<Property Name="Emails" Type="Collection(Edm.String)"/>
<Property Name="AddressInfo" Type="Collection(Trippin.Model.Location)"/>
<Property Name="HomeAddress" Type="Trippin.Model.Location"/>
<Property Name="FavoriteFeature" Type="Trippin.Model.Feature" Nullable="false"/>
<Property Name="Features" Type="Collection(Trippin.Model.Feature)" Nullable="false"/>
<NavigationProperty Name="Friends" Type="Collection(Trippin.Model.Person)"/>
<NavigationProperty Name="BestFriend" Type="Trippin.Model.Person"/>
<NavigationProperty Name="Trips" Type="Collection(Trippin.Model.Trip)"/>
</EntityType>
<EntityType Name="Airline">
<Key>
<PropertyRef Name="AirlineCode"/>
</Key>
<Property Name="AirlineCode" Type="Edm.String" Nullable="false"/>
<Property Name="Name" Type="Edm.String"/>
</EntityType>
<EntityType Name="Airport">
<Key>
<PropertyRef Name="IcaoCode"/>
</Key>
<Property Name="Name" Type="Edm.String"/>
<Property Name="IcaoCode" Type="Edm.String" Nullable="false"/>
<Property Name="IataCode" Type="Edm.String"/>
<Property Name="Location" Type="Trippin.Model.AirportLocation"/>
</EntityType>
<ComplexType Name="Location">
<Property Name="Address" Type="Edm.String"/>
<Property Name="City" Type="Trippin.Model.City"/>
</ComplexType>
<ComplexType Name="City">
<Property Name="Name" Type="Edm.String"/>
<Property Name="CountryRegion" Type="Edm.String"/>
<Property Name="Region" Type="Edm.String"/>
</ComplexType>
<ComplexType Name="AirportLocation" BaseType="Trippin.Model.Location">
<Property Name="Loc" Type="Edm.GeographyPoint"/>
</ComplexType>
<ComplexType Name="EventLocation" BaseType="Trippin.Model.Location">
<Property Name="BuildingInfo" Type="Edm.String"/>
</ComplexType>
<EntityType Name="Trip">
<Key>
<PropertyRef Name="TripId"/>
</Key>
<Property Name="TripId" Type="Edm.Int32" Nullable="false"/>
<Property Name="ShareId" Type="Edm.Guid" Nullable="false"/>
<Property Name="Name" Type="Edm.String"/>
<Property Name="Budget" Type="Edm.Single" Nullable="false"/>
<Property Name="Description" Type="Edm.String"/>
<Property Name="Tags" Type="Collection(Edm.String)"/>
<Property Name="StartsAt" Type="Edm.DateTimeOffset" Nullable="false"/>
<Property Name="EndsAt" Type="Edm.DateTimeOffset" Nullable="false"/>
<NavigationProperty Name="PlanItems" Type="Collection(Trippin.Model.PlanItem)"/>
</EntityType>
<EntityType Name="PlanItem">
<Key>
<PropertyRef Name="PlanItemId"/>
</Key>
<Property Name="PlanItemId" Type="Edm.Int32" Nullable="false"/>
<Property Name="ConfirmationCode" Type="Edm.String"/>
<Property Name="StartsAt" Type="Edm.DateTimeOffset" Nullable="false"/>
<Property Name="EndsAt" Type="Edm.DateTimeOffset" Nullable="false"/>
<Property Name="Duration" Type="Edm.Duration" Nullable="false"/>
</EntityType>
<EntityType Name="Event" BaseType="Trippin.Model.PlanItem">
<Property Name="OccursAt" Type="Trippin.Model.EventLocation"/>
<Property Name="Description" Type="Edm.String"/>
</EntityType>
<EntityType Name="PublicTransportation" BaseType="Trippin.Model.PlanItem">
<Property Name="SeatNumber" Type="Edm.String"/>
</EntityType>
<EntityType Name="Flight" BaseType="Trippin.Model.PublicTransportation">
<Property Name="FlightNumber" Type="Edm.String"/>
<NavigationProperty Name="Airline" Type="Trippin.Model.Airline"/>
<NavigationProperty Name="From" Type="Trippin.Model.Airport"/>
<NavigationProperty Name="To" Type="Trippin.Model.Airport"/>
</EntityType>
<EntityType Name="Employee" BaseType="Trippin.Model.Person">
<Property Name="Cost" Type="Edm.Int64" Nullable="false"/>
<NavigationProperty Name="Peers" Type="Collection(Trippin.Model.Person)"/>
</EntityType>
<EntityType Name="Manager" BaseType="Trippin.Model.Person">
<Property Name="Budget" Type="Edm.Int64" Nullable="false"/>
<Property Name="BossOffice" Type="Trippin.Model.Location"/>
<NavigationProperty Name="DirectReports" Type="Collection(Trippin.Model.Person)"/>
</EntityType>
<EnumType Name="PersonGender">
<Member Name="Male" Value="0"/>
<Member Name="Female" Value="1"/>
<Member Name="Unknown" Value="2"/>
</EnumType>
<EnumType Name="Feature">
<Member Name="Feature1" Value="0"/>
<Member Name="Feature2" Value="1"/>
<Member Name="Feature3" Value="2"/>
<Member Name="Feature4" Value="3"/>
</EnumType>
<Function Name="GetPersonWithMostFriends">
<ReturnType Type="Trippin.Model.Person"/>
</Function>
<Function Name="GetNearestAirport">
<Parameter Name="lat" Type="Edm.Double" Nullable="false"/>
<Parameter Name="lon" Type="Edm.Double" Nullable="false"/>
<ReturnType Type="Trippin.Model.Airport"/>
</Function>
<Function Name="GetFavoriteAirline" IsBound="true" EntitySetPath="person">
<Parameter Name="person" Type="Trippin.Model.Person"/>
<ReturnType Type="Trippin.Model.Airline"/>
</Function>
<Function Name="GetFriendsTrips" IsBound="true">
<Parameter Name="person" Type="Trippin.Model.Person"/>
<Parameter Name="userName" Type="Edm.String" Nullable="false" Unicode="false"/>
<ReturnType Type="Collection(Trippin.Model.Trip)"/>
</Function>
<Function Name="GetInvolvedPeople" IsBound="true">
<Parameter Name="trip" Type="Trippin.Model.Trip"/>
<ReturnType Type="Collection(Trippin.Model.Person)"/>
</Function>
<Action Name="ResetDataSource"/>
<Action Name="UpdateLastName" IsBound="true">
<Parameter Name="person" Type="Trippin.Model.Person"/>
<Parameter Name="lastName" Type="Edm.String" Nullable="false" Unicode="false"/>
<ReturnType Type="Edm.Boolean" Nullable="false"/>
</Action>
<Action Name="ShareTrip" IsBound="true">
<Parameter Name="personInstance" Type="Trippin.Model.Person"/>
<Parameter Name="userName" Type="Edm.String" Nullable="false" Unicode="false"/>
<Parameter Name="tripId" Type="Edm.Int32" Nullable="false"/>
</Action>
</Schema>
<Schema xmlns="http://docs.oasis-open.org/odata/ns/edm" Namespace="Trippin.Data">
<EntityContainer Name="Container">
<EntitySet Name="People" EntityType="Trippin.Model.Person">
<NavigationPropertyBinding Path="BestFriend" Target="People"/>
<NavigationPropertyBinding Path="Friends" Target="People"/>
<NavigationPropertyBinding Path="Trippin.Model.Employee/Peers" Target="People"/>
<NavigationPropertyBinding Path="Trippin.Model.Manager/DirectReports" Target="People"/>
</EntitySet>
<EntitySet Name="Airlines" EntityType="Trippin.Model.Airline">
<Annotation Term="Org.OData.Core.V1.OptimisticConcurrency">
<Collection>
<PropertyPath>Name</PropertyPath>
</Collection>
</Annotation>
</EntitySet>
<EntitySet Name="Airports" EntityType="Trippin.Model.Airport"/>
<Singleton Name="Me" Type="Trippin.Model.Person">
<NavigationPropertyBinding Path="BestFriend" Target="People"/>
<NavigationPropertyBinding Path="Friends" Target="People"/>
<NavigationPropertyBinding Path="Trippin.Model.Employee/Peers" Target="People"/>
<NavigationPropertyBinding Path="Trippin.Model.Manager/DirectReports" Target="People"/>
</Singleton>
<FunctionImport Name="GetPersonWithMostFriends" Function="Trippin.Model.GetPersonWithMostFriends" EntitySet="People"/>
<FunctionImport Name="GetNearestAirport" Function="Trippin.Model.GetNearestAirport" EntitySet="Airports"/>
<ActionImport Name="ResetDataSource" Action="Trippin.Model.ResetDataSource"/>
</EntityContainer>
</Schema>
</edmx:DataServices>
</edmx:Edmx>`

var parsedMultiSchemaEdmx edmxDataServices
var multiSchemaEdmxParseError error
var hasParsedMultiSchemaEdmx = false

func getParsedMultiSchemaEdmx() (edmxDataServices, error) {
	if !hasParsedMultiSchemaEdmx {
		parsedMultiSchemaEdmx, multiSchemaEdmxParseError = parseEdmx([]byte(multiSchemaEdmxSchema))
		hasParsedMultiSchemaEdmx = true
	}
	return parsedMultiSchemaEdmx, multiSchemaEdmxParseError
}

func Test_Parse_multi_schema_edmx(t *testing.T) {
	ds, err := getParsedMultiSchemaEdmx()
	assert.NoError(t, err)
	assert.Equal(t, "Trippin.Model", ds.Schemas["Trippin.Model"].Namespace)

	edmx := ds.Schemas["Trippin.Model"]
	edmx2 := ds.Schemas["Trippin.Data"]

	personEntityType, ok := edmx.EntityTypes["Person"]
	assert.True(t, ok)
	assert.Equal(t, "Person", personEntityType.Name)
	assert.Len(t, personEntityType.Properties, 11)
	usernameProperty, ok := personEntityType.Properties["UserName"]
	assert.True(t, ok)
	assert.Equal(t, "Edm.String", usernameProperty.Type)
	assert.Equal(t, "string", usernameProperty.goType())
	lastNameProperty, ok := personEntityType.Properties["LastName"]
	assert.True(t, ok)
	assert.Equal(t, "Edm.String", lastNameProperty.Type)
	assert.Equal(t, "nullable.Nullable[string]", lastNameProperty.goType())

	peopleEntitySet := edmx2.EntitySets["People"]
	assert.Equal(t, "Trippin.Model.Person", peopleEntitySet.EntityType)
	assert.Equal(t, personEntityType, peopleEntitySet.getEntityType())
}
